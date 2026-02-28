package zapKit

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewSimpleConsoleLogger() *zap.Logger {
	return NewLogger(nil)
}

// NewLogger
/*
PS:
(1) 自定义字段(Field)，创建 core 和 logger 时都能添加.

@param core		可以为nil
@param options	可以不传

e.g. case: core传nil，options不传
	(1) [Encoder] 人类可读的多行输出
	(2) [Encoder] 时间格式: "2024-06-28T09:15:16.176+0800"
	(3) [Encoder] 日志级别大写且有颜色(color)
	(4) [Encoder] Message字段无前缀
	(5) [Encoder] caller 字段左对齐，最小长度为 30

	(6) [Core] 仅有1个输出: 输出到控制台(并发安全地输出到os.Stdout)
	(7) [Core] 仅有1个输出: 日志级别(level)为 DEBUG

	(8) [Logger] 有 Caller 且 CallerSkip == 0
	(9) [Logger] Development == false，即生产模式
	(10) [Logger] ErrorOutput 使用默认值: 并发安全地输出到os.Stderr
	(11) [Logger] DPanicLevel 及以上级别的日志输出，会附带堆栈信息(stack trace)
*/
func NewLogger(core zapcore.Core, options ...LoggerOption) (logger *zap.Logger) {
	if core == nil {
		core = NewCore(NewEncoder(), os.Stdout, zapcore.DebugLevel)
	}

	opts := loadOptions(options...)

	var zapOptions []zap.Option
	// Development
	if opts.Development {
		zapOptions = append(zapOptions, zap.Development())
	}
	// ErrorOutput
	zapOptions = append(zapOptions, zap.ErrorOutput(opts.ErrorOutput))
	// Caller
	zapOptions = append(zapOptions, zap.WithCaller(opts.Caller))
	// CallerSkip
	zapOptions = append(zapOptions, zap.AddCallerSkip(opts.CallerSkip))
	// AddStacktrace
	if opts.AddStacktrace != nil {
		zapOptions = append(zapOptions, zap.AddStacktrace(opts.AddStacktrace))
	}
	// Clock
	if opts.Clock != nil {
		zapOptions = append(zapOptions, zap.WithClock(opts.Clock))
	}
	// Fields
	if len(opts.Fields) > 0 {
		zapOptions = append(zapOptions, zap.Fields(opts.Fields...))
	}
	// PanicHook
	if opts.PanicHook != nil {
		zapOptions = append(zapOptions, zap.WithPanicHook(opts.PanicHook))
	}
	// FatalHook
	if opts.FatalHook != nil {
		zapOptions = append(zapOptions, zap.WithFatalHook(opts.FatalHook))
	}

	logger = zap.New(core, zapOptions...)
	return
}

func NewSugarLogger(core zapcore.Core, options ...LoggerOption) *zap.SugaredLogger {
	return NewLogger(core, options...).Sugar()
}
