package nacosKit

import (
	"net/url"

	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/richelieu042/chimera/v3/src/copyKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/intKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/netKit"
	"github.com/richelieu042/chimera/v3/src/urlKit"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

var (
	clientConfig  *constant.ClientConfig
	serverConfigs []constant.ServerConfig
)

func MustSetUp(config *Config, options ...constant.ClientOption) {
	err := SetUp(config, options...)
	if err != nil {
		console.Fatalf("Fail to set up, error: %+v", err)
	}
}

// SetUp
/*
@param options		!!!:
					(1) 建议配置 客户端的缓存目录（default is current path）		constant.WithCacheDir
					(2) 建议配置 客户端的日志目录（default is current path）		constant.WithLogDir
					(3) 建议配置 客户端的日志级别（default value is info）		constant.WithLogLevel	"debug" || "info"（默认）...
*/
func SetUp(config *Config, options ...constant.ClientOption) (err error) {
	defer func() {
		if err != nil {
			clientConfig = nil
			serverConfigs = nil
		}
	}()

	/* (0) validate */
	if err = validateKit.Struct(config); err != nil {
		err = errorKit.Wrapf(err, "Fail to verify")
		return
	}

	/* (1) clientConfig */
	options1 := []constant.ClientOption{
		constant.WithNamespaceId(config.NamespaceId),
		constant.WithBeatInterval(5000), // 5000ms
		constant.WithNotLoadCacheAtStart(true),
	}
	options1 = append(options1, options...)
	clientConfig = constant.NewClientConfig(options1...)

	/* (2) serverConfigs */
	if err = sliceKit.AssertNotEmpty(config.Addrs, "config.Addrs"); err != nil {
		return
	}
	for _, addr := range config.Addrs {
		if strKit.IsBlank(addr) {
			continue
		}

		var u *url.URL
		u, err = urlKit.Parse(addr)
		if err != nil {
			err = errorKit.Wrapf(err, "Fail to parse address(%s).", addr)
			return
		}
		var port uint64
		port, err = intKit.ToUint64E(u.Port())
		if err != nil {
			err = errorKit.Wrapf(err, "Invalid address(%s) with port string(%s).", addr, u.Port())
			return
		}
		if err = netKit.AssertValidPort(int(port)); err != nil {
			err = errorKit.Wrapf(err, "Invalid address(%s) with port(%d).", addr, port)
			return
		}

		serverConfigs = append(serverConfigs, constant.ServerConfig{
			Scheme:      u.Scheme,
			IpAddr:      u.Hostname(),
			Port:        port,
			ContextPath: u.Path,
		})
	}
	if sliceKit.IsEmpty(serverConfigs) {
		err = errorKit.Newf("No valid address.")
		return
	}

	/* (3) test（以防配置有问题） */
	cc, err := NewConfigClient()
	if err != nil {
		return
	}
	defer cc.CloseClient()
	nc, err := NewNamingClient()
	if err != nil {
		return
	}
	defer nc.CloseClient()

	return
}

func GetClientConfigCopy(options ...constant.ClientOption) (*constant.ClientConfig, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, NotSetUpError
	}

	// 深拷贝
	clientConfig1 := copyKit.DeepCopy(clientConfig)

	// 再次修改 *constant.ClientConfig
	for _, option := range options {
		option(clientConfig1)
	}

	return clientConfig, nil
}

// GetNamespaceId
/*
PS: 如果一个服务仅使用一个NamespaceId，可以调用此方法.
*/
func GetNamespaceId() (string, error) {
	if clientConfig == nil || serverConfigs == nil {
		return "", NotSetUpError
	}

	return clientConfig.NamespaceId, nil
}
