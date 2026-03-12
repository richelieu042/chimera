package zapKit

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

const (
	// OutputFormatConsole 人类可读的多行输出
	OutputFormatConsole outputFormat = iota

	// OutputFormatJson JSON格式输出
	OutputFormatJson
)

type (
	outputFormat uint8

	encoderOptions struct {
		// OutputFormat 输出类型
		OutputFormat outputFormat

		MessagePrefix string

		EncodeLevel  zapcore.LevelEncoder
		EncodeTime   zapcore.TimeEncoder
		EncodeCaller zapcore.CallerEncoder
	}

	EncoderOption func(opts *encoderOptions)
)

func (opts *encoderOptions) IsOutputFormatConsole() bool {
	if opts == nil {
		return true
	}
	return opts.OutputFormat == OutputFormatConsole
}

func loadEncoderOptions(options ...EncoderOption) *encoderOptions {
	opts := &encoderOptions{
		OutputFormat:  OutputFormatConsole,
		MessagePrefix: "",
		EncodeLevel:   nil,
		EncodeTime:    nil,
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			// 让 caller 字段左对齐，提高可读性，参考：https://claude.ai/share/1358e58e-27f2-4305-b2fa-7c7fc8f6ae4a
			s := caller.TrimmedPath()
			if len(s) < minCallerLen {
				s = s + strings.Repeat(" ", minCallerLen-len(s))
			}
			enc.AppendString(s)
		},
	}

	for _, option := range options {
		option(opts)
	}

	// EncodeLevel
	if opts.EncodeLevel == nil {
		if opts.IsOutputFormatConsole() {
			// 大写 && 带颜色
			opts.EncodeLevel = zapcore.CapitalColorLevelEncoder
		} else {
			// 大写
			opts.EncodeLevel = zapcore.CapitalLevelEncoder
		}
	}
	// EncodeTime
	if opts.EncodeTime == nil {
		if opts.IsOutputFormatConsole() {
			opts.EncodeTime = zapcore.ISO8601TimeEncoder
		} else {
			opts.EncodeTime = zapcore.EpochTimeEncoder
		}
	}

	return opts
}

func WithEncoderOutputFormatConsole() EncoderOption {
	return func(opts *encoderOptions) {
		opts.OutputFormat = OutputFormatConsole
	}
}

func WithEncoderOutputFormatJson() EncoderOption {
	return func(opts *encoderOptions) {
		opts.OutputFormat = OutputFormatJson
	}
}

func WithEncoderMessagePrefix(prefix string) EncoderOption {
	return func(opts *encoderOptions) {
		opts.MessagePrefix = prefix
	}
}

func WithEncoderEncodeLevel(encodeLevel zapcore.LevelEncoder) EncoderOption {
	return func(opts *encoderOptions) {
		opts.EncodeLevel = encodeLevel
	}
}

func WithEncoderEncodeTime(encodeTime zapcore.TimeEncoder) EncoderOption {
	return func(opts *encoderOptions) {
		opts.EncodeTime = encodeTime
	}
}
