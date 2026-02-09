package etcdKit

import (
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
)

type (
	Config struct {
		Endpoints []string `json:"endpoints" yaml:"endpoints"`
	}
)

func (config *Config) Check() error {
	if err := interfaceKit.AssertNotNil(config, "config"); err != nil {
		return err
	}

	config.Endpoints = sliceKit.PolyfillStringSlice(config.Endpoints)
	if err := sliceKit.AssertNotEmpty(config.Endpoints, "config.Endpoints"); err != nil {
		return err
	}

	return nil
}
