package nacosKit

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/idKit"
)

// NewConfigClient 创建 动态配置(config) 客户端.
/*
!!!:
(1) 需要先set up.
(2) config_client.IConfigClient 实例，如果不用了需要"手动关闭".

@param options 可以用于修改: NamespaceId、CacheDir、LogDir、LogLevel...
*/
func NewConfigClient(options ...constant.ClientOption) (config_client.IConfigClient, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, NotSetUpError
	}

	clientConfig1, err := GetClientConfigCopy(options...)
	if err != nil {
		return nil, err
	}
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  clientConfig1,
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		return nil, err
	}

	/* verify */
	tmp := fmt.Sprintf("%s_%s", consts.ProjectName, idKit.NewULID())
	_, err = client.GetConfig(vo.ConfigParam{
		DataId: tmp,
		Group:  tmp,
	})
	if err != nil {
		err = errorKit.Wrapf(err, "Fail to pass verification, check the configuration please!")
		return nil, err
	}

	return client, nil
}

// NewNamingClient 创建 服务发现(naming) 客户端.
/*
!!!:
(1) 需要先set up;
(2) naming_client.INamingClient 实例，如果不用了需要"手动关闭";
(3) !!!: 不要在 Nacos管理页面 上修改服务实例的信息（weight等），否则可能导致后续 RegisterInstance 、UpdateInstance 会失败（虽然方法返回true, nil）.

@param options 可以用于修改: NamespaceId...
*/
func NewNamingClient(options ...constant.ClientOption) (naming_client.INamingClient, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, NotSetUpError
	}

	clientConfig1, err := GetClientConfigCopy(options...)
	if err != nil {
		return nil, err
	}
	client, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  clientConfig1,
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		return nil, err
	}

	/* verify */
	serviceName := fmt.Sprintf("%s_%s", consts.ProjectName, idKit.NewULID())
	_, err = client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
	})
	if err != nil {
		// TODO: 此处比较low，比较错误的文本内容，看后续库有没有更新吧.
		tmp := "instance list is empty!"
		if !strKit.EqualsIgnoreCase(err.Error(), tmp) {
			err = errorKit.Wrapf(err, "Fail to pass verification, check the configuration please!")
			return nil, err
		}
	}

	return client, nil
}
