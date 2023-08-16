package demo

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rshulabs/micro-frame/internal/demo/config"
	"github.com/rshulabs/micro-frame/internal/demo/controller"
	"github.com/rshulabs/micro-frame/internal/demo/store"
)

// 运行函数
func Run(cfg *config.Config) error {
	http := controller.NewHttpService(cfg)
	srv := store.NewGrpcService(cfg)
	if cfg.App.IsStartHttp {
		go http.Start()
	}
	go srv.Start()
	ch := make(chan os.Signal, 1)
	defer close(ch)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	select {
	case <-ch:
		http.WithStop()
		if err := srv.Stop(); err != nil {
			return err
		}
	}
	return nil
}
