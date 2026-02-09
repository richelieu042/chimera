package logrusKit

import (
	"github.com/richelieu042/chimera/v3/src/core/ioKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/sirupsen/logrus"
)

// NewLogger
/*
PS:
(1) 默认输出到 控制台(os.Stderr);
(2) 如果希望 输出到文件（可选是否rotatable），可以使用 WithOutput()，详见下例.

@param options 可以什么都不配置（此时输出到控制台）
*/
func NewLogger(options ...LoggerOption) *logrus.Logger {
	opts := loadOptions(options...)

	logger := logrus.New()
	logger.SetFormatter(opts.formatter)
	logger.SetReportCaller(opts.reportCaller)
	logger.SetLevel(opts.level)
	if opts.output != nil {
		logger.SetOutput(opts.output)
	}

	// msgPrefix
	if strKit.IsNotEmpty(opts.msgPrefix) {
		hook := &prefixHook{prefix: opts.msgPrefix}
		logger.AddHook(hook)
	}

	if opts.disableQuote {
		DisableQuote(logger)
	}

	return logger
}

// DisposeLogger 释放资源（主要针对文件日志）
func DisposeLogger(logger *logrus.Logger) error {
	if logger == nil {
		return nil
	}

	return ioKit.TryToClose(logger.Out)
}
