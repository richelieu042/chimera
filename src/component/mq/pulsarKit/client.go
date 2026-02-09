package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/log/logrusKit"
	"github.com/richelieu042/chimera/v3/src/netKit"
	"github.com/sirupsen/logrus"
)

// NewClient
/*
@param logPath 客户端的日志输出（为空则输出到控制台; 不会rotate）
*/
func NewClient(addresses []string, logPath string) (pulsar.Client, error) {
	/* url */
	addresses, err := netKit.PolyfillHosts(addresses, 1)
	if err != nil {
		return nil, err
	}
	url := UrlPrefix + sliceKit.Join(addresses, ",")

	//tmp, err := netKit.ProcessAddresses(addresses)
	//if err != nil {
	//	return nil, err
	//}
	//url := UrlPrefix + sliceKit.Join(tmp, ",")

	/* logger */
	var logger log.Logger
	if strKit.IsNotEmpty(logPath) {
		fileLogger, err := logrusKit.NewFileLogger(logPath)
		if err != nil {
			return nil, err
		}

		// 文件日志的日志级别 与 logrus全局日志级别 保持一致
		fileLogger.SetLevel(logrus.GetLevel())

		logger = log.NewLoggerWithLogrus(fileLogger)
	}

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:    url,
		Logger: logger,
	})
	if err != nil {
		err = errorKit.Wrapf(err, "fail to new client")
		return nil, err
	}

	return client, nil
}
