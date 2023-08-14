package options

import (
	"encoding/json"

	"github.com/rshulabs/micro-frame/internal/pkg/options"
	"github.com/rshulabs/micro-frame/pkg/app"
)

type Options struct {
	App  *options.AppOption  `json:"app" mapstructure:"app" yaml:"app"`
	Grpc *options.GrpcOption `json:"grpc" mapstructure:"grpc" yaml:"grpc"`
	Http *options.HttpOption `json:"http" mapstructure:"http" yaml:"http"`
}

func NewOptions() *Options {
	return &Options{
		App:  options.NewAppOption(),
		Grpc: options.NewGrpcOption(),
		Http: options.NewHttpOption(),
	}
}

func (o *Options) Flags() (fss app.FlagSets) {
	// 分组
	o.App.AddFlags(fss.FlagSet("app"))
	o.Grpc.AddFlags(fss.FlagSet("grpc"))
	o.Http.AddFlags(fss.FlagSet("http"))
	return fss
}

func (o *Options) Complete() error {
	return nil
}

func (o *Options) String() string {
	bytes, _ := json.Marshal(o)
	return string(bytes)
}
