package logrusKit

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type (
	loggerOptions struct {
		// formatter 日志格式
		formatter logrus.Formatter

		// reportCaller 默认: true
		reportCaller bool

		// level 日志级别，默认: logrus.DebugLevel
		level logrus.Level

		// output 默认: os.Stderr
		output io.Writer

		// msgPrefix
		/*
			日志输出的msg属性的前缀，
			(1) 默认: ""
			(2) 非空的话，会拼接在msg前面（但并不会在中间加上" "）
		*/
		msgPrefix string

		// disableQuote 默认: false
		disableQuote bool
	}

	LoggerOption func(opts *loggerOptions)
)

func loadOptions(options ...LoggerOption) *loggerOptions {
	/* 默认值s */
	opts := &loggerOptions{
		formatter:    nil,
		reportCaller: true,
		level:        logrus.DebugLevel,
		// 默认: 输出到控制台
		output:       os.Stderr,
		msgPrefix:    "",
		disableQuote: false,
	}

	for _, option := range options {
		option(opts)
	}

	// 容错，以防"调用方"瞎传
	if opts.formatter == nil {
		opts.formatter = NewDefaultTextFormatter()
	}
	return opts
}

func WithFormatter(formatter logrus.Formatter) LoggerOption {
	return func(opts *loggerOptions) {
		opts.formatter = formatter
	}
}

func WithReportCaller(reportCaller bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.reportCaller = reportCaller
	}
}

func WithLevel(level logrus.Level) LoggerOption {
	return func(opts *loggerOptions) {
		opts.level = level
	}
}

func WithOutput(output io.Writer) LoggerOption {
	return func(opts *loggerOptions) {
		opts.output = output
	}
}

func WithMsgPrefix(msgPrefix string) LoggerOption {
	return func(opts *loggerOptions) {
		opts.msgPrefix = msgPrefix
	}
}

func WithDisableQuote(disableQuote bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.disableQuote = disableQuote
	}
}
