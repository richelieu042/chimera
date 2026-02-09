package consulKit

import (
	"github.com/hashicorp/consul/api"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
)

// NewClient
/*
@param config 	(1) 不能为 nil;
				(2) Address 默认为: "127.0.0.1:8500"，Scheme 默认为: "http";
				(3) 一般情况下，配置下 Address 即可.
*/
func NewClient(config *api.Config) (*api.Client, error) {
	if err := interfaceKit.AssertNotNil(config, "config"); err != nil {
		return nil, err
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	// 防止 Address 是无效的
	_, err = client.Agent().Checks()
	if err != nil {
		return nil, errorKit.Wrapf(err, "fail to get the locally registered checks")
	}

	return client, nil
}
