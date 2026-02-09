package etcdKit

import (
	"io"
	"sync"
	"time"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/log/zapKit"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var client *clientv3.Client
var setupOnce sync.Once

func MustSetUp(config *Config, logPath string) {
	if err := setUp(config, logPath); err != nil {
		console.Fatalf("Fail to set up, error: %s", err.Error())
	}
}

// setUp
/*
TODO: 可以参考 go-zero 中 registry.go 的 internal.DialClient.

PS:
(1) 如果 Endpoints 无效，会返回error(context.DeadlineExceeded).

@param logPath 为空则输出到控制台
*/
func setUp(config *Config, logPath string) (err error) {
	if err = config.Check(); err != nil {
		return
	}

	setupOnce.Do(func() {
		/* etcd客户端日志输出 */
		var logger *zap.Logger
		if strKit.IsNotEmpty(logPath) {
			var writer io.Writer
			writer, err = fileKit.CreateInAppendMode(logPath)
			if err != nil {
				return
			}

			// logger日志级别 等同于 logrus的全局日志级别
			var level zapcore.Level
			level, err = zapKit.StringToLevel(logrus.GetLevel().String())
			if err != nil {
				return
			}

			enc := zapKit.NewEncoder()
			core := zapKit.NewCore(enc, zapKit.NewWriteSyncer(writer), level)
			logger = zapKit.NewLogger(core)
		}

		v3Config := clientv3.Config{
			Endpoints: config.Endpoints,
			Logger:    logger,

			AutoSyncInterval:     time.Minute,
			DialTimeout:          time.Second * 5,
			DialKeepAliveTime:    time.Second * 5,
			DialKeepAliveTimeout: time.Second * 5,
			RejectOldCluster:     true,
			PermitWithoutStream:  true,
		}
		client, err = clientv3.New(v3Config)
	})
	return
}

// GetClient
/*
PS:
(1) 要使用 KV 的情况下，建议调用 clientv3.NewKV() 以实例化一个用于操作etcd的KV（内置错误重试机制）.
(2) 租约相关需要用到 *clientv3.Client实例.
*/
func GetClient() (*clientv3.Client, error) {
	if client == nil {
		return nil, NotSetupError
	}
	return client, nil
}
