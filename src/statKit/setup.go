package statKit

import (
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/cronKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func MustSetup(logPath string) {
	if err := Setup(logPath); err != nil {
		console.Fatalf("Fail to set up, error: %+v", err)
	}
}

func Setup(logPath string) error {
	if strKit.IsBlank(logPath) {
		// (1) 输出到: 控制台
		logger = zapKit.NewSimpleConsoleLogger().Sugar()
	} else {
		// (2) 输出到: 文件日志
		enc := zapKit.NewEncoder()
		f, err := fileKit.CreateInAppendMode(logPath)
		if err != nil {
			return err
		}
		ws := zapKit.NewLockedWriteSyncer(f)
		core := zapKit.NewCore(enc, ws, zap.DebugLevel)
		logger = zapKit.NewLogger(core).Sugar()
	}

	c, _, err := cronKit.NewCronWithTask("@every 30s", func() {
		PrintStats(logger)
	})
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
