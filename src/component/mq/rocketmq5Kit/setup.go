package rocketmq5Kit

import (
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

var (
	NotSetupError = errorKit.Newf("haven’t been set up correctly")
)

var config *Config

func MustSetUp(c *Config, clientLogPath string, verifyConfig *VerifyConfig) {
	if err := SetUp(c, clientLogPath, verifyConfig); err != nil {
		console.Fatalf("Fail to set up, error: %s", err.Error())
	}
}

// SetUp
/*
@param clientLogPath	可以为""（输出到控制台）
@param verifyConfig		可以为nil（不进行验证）
*/
func SetUp(c *Config, clientLogPath string, verifyConfig *VerifyConfig) (err error) {
	defer func() {
		if err != nil {
			config = nil
		}
	}()

	if err = validateKit.Struct(c); err != nil {
		return
	}
	// Richelieu: 防止客户端库源码内部panic
	if c.Credentials == nil {
		c.Credentials = &credentials.SessionCredentials{
			AccessKey:     "",
			AccessSecret:  "",
			SecurityToken: "",
		}
	}

	// 客户端日志输出
	if err = setClientLog(clientLogPath); err != nil {
		return
	}

	config = c

	// verify
	if err = verify(verifyConfig); err != nil {
		return
	}

	return
}
