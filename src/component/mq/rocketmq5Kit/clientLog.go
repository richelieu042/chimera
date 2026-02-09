package rocketmq5Kit

import (
	"os"

	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu042/chimera/v3/src/core/envKit"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

func setClientLog(path string) error {
	if strKit.IsEmpty(path) {
		// (1) 输出到控制台
		if err := os.Setenv(rmq_client.ENABLE_CONSOLE_APPENDER, "true"); err != nil {
			return err
		}
	} else {
		// (2) 输出到日志文件
		if err := fileKit.AssertNotExistOrIsFile(path); err != nil {
			return err
		}
		if err := fileKit.MkParentDirs(path); err != nil {
			return err
		}

		dir := pathKit.ParentDir(path)
		name := fileKit.GetFileName(path)

		if err := envKit.SetEnvs(map[string]string{
			rmq_client.ENABLE_CONSOLE_APPENDER: "false",
			rmq_client.CLIENT_LOG_ROOT:         dir,
			rmq_client.CLIENT_LOG_FILENAME:     name,
		}); err != nil {
			return err
		}
	}

	rmq_client.ResetLogger()
	return nil
}
