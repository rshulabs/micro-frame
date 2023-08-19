package discovery

// 注册抽象
type Registry interface {
	// 注册
	Regisrty(srv Service) error
	// 注销
	DeRegistry(srv Service) error
}

// 服务类型
type Service interface {
	Name() string
	Addr() string
}

// 发现抽象
type Discovery interface {
	// 获取服务器的一个地址
	GetServiceAddr(srv *DisService) (string, error)
	// 监控服务地址变化
	WatchService(srv *DisService) error
}

type DisService struct {
	ServiceName string
	Replicas    int
	Url         string
	Endpoints   []string
}

type Option func(*DisService)

func WithReplicas(replicas int) Option {
	return func(o *DisService) {
		o.Replicas = replicas
	}
}

func NewDisService(srvName string, url string, opts ...Option) *DisService {
	if srvName == "" && url == "" {
		return nil
	}
	s := &DisService{
		Replicas:  100,
		Endpoints: []string{"192.168.60.120:2379"},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
