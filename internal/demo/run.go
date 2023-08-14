package demo

import (
	"github.com/rshulabs/micro-frame/internal/demo/config"
	"github.com/rshulabs/micro-frame/pkg/logx"
)

// 运行函数
func Run(cfg *config.Config) error {
	logx.Info(cfg.String())
	return nil
}
