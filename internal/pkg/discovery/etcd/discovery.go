package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/rshulabs/micro-frame/internal/pkg/discovery"
	lb "github.com/rshulabs/micro-frame/internal/pkg/discovery/lb/consisthash"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdDiscovery struct {
	cli *clientv3.Client
}

func NewEtcdDiscovery(endpoints []string) (*EtcdDiscovery, error) {
	// 校验
	if len(endpoints) == 0 {
		return nil, fmt.Errorf("endpoints cannot be empty")
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &EtcdDiscovery{
		cli: cli,
	}, nil
}

func (d *EtcdDiscovery) GetServiceAddr(srv *discovery.DisService) (string, error) {
	// get --prefix
	gResp, err := d.cli.Get(context.Background(), srv.ServiceName, clientv3.WithPrefix())
	if err != nil {
		return "", err
	}
	if len(gResp.Kvs) == 0 {
		return "", fmt.Errorf("%s service is not found", srv.ServiceName)
	}
	m := lb.NewMap(srv.Replicas, nil)
	for _, v := range gResp.Kvs {
		go m.Add(v.String())
	}
	// 采用随机LB，做负载均衡
	// randIndex := rand.Intn(len(gResp.Kvs)) // [0,n)
	// addr := string(gResp.Kvs[randIndex].Value)
	addr := m.Get(srv.Url)
	return addr, nil
}

func (d *EtcdDiscovery) WatchService(srv discovery.DisService) error {
	return nil
}
