package demo

import (
	"github.com/rshulabs/micro-frame/internal/demo/config"
	"github.com/rshulabs/micro-frame/pkg/logx"
)

func Run(cfg *config.Config) error {
	logx.Info(cfg.String())
	return nil
}
