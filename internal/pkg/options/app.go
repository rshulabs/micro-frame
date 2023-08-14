package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

type AppOption struct {
	Name string `json:"name" mapstructure:"name" yaml:"name"`
}

func NewAppOption() *AppOption {
	return &AppOption{
		Name: "",
	}
}

func (a *AppOption) Validate() []error {
	var errors []error
	if a.Name == "" {
		errors = append(errors, fmt.Errorf("app name is required"))
	}
	return errors
}

func (a *AppOption) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&a.Name, "app.name", a.Name, "The binary name of app")
}
