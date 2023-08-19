package controller

import (
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"github.com/rshulabs/micro-frame/internal/demo/store/pb"
	"github.com/rshulabs/micro-frame/internal/pkg/discovery"
	"github.com/rshulabs/micro-frame/internal/pkg/discovery/etcd"
	"google.golang.org/grpc"
)

type Service struct {
	Srv pb.DemoClient
}

func NewService(srv *discovery.DisService) *Service {
	discover, err := etcd.NewEtcdDiscovery(srv.Endpoints)
	if err != nil {
		panic(err)
	}
	addr, err := discover.GetServiceAddr(srv)
	if err != nil {
		panic(err)
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	// 重试
	opts = append(opts, grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor()))
	// 超时
	opts = append(opts, grpc.WithUnaryInterceptor(timeout.UnaryClientInterceptor(1*time.Second)))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(err)
	}
	cli := pb.NewDemoClient(conn)
	return &Service{
		Srv: cli,
	}
}
