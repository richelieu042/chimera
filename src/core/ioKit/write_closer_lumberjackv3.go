package ioKit

import (
	"time"

	"github.com/natefinch/lumberjack/v3"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/mathKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

type (
	lumberjackOptions struct {
		maxAge     time.Duration
		maxBackups int
		localTime  bool
		compress   bool
	}

	LumberJackOption func(opts *lumberjackOptions)
)

func WithMaxAge(maxAge time.Duration) LumberJackOption {
	return func(opts *lumberjackOptions) {
		opts.maxAge = maxAge
	}
}

func WithMaxBackups(maxBackups int) LumberJackOption {
	return func(opts *lumberjackOptions) {
		opts.maxBackups = maxBackups
	}
}

func WithLocalTime(localTime bool) LumberJackOption {
	return func(opts *lumberjackOptions) {
		opts.localTime = localTime
	}
}

func WithCompress(compress bool) LumberJackOption {
	return func(opts *lumberjackOptions) {
		opts.compress = compress
	}
}

func loadOptions(options ...LumberJackOption) *lumberjackOptions {
	opts := &lumberjackOptions{
		maxAge:     0,
		maxBackups: 0,
		localTime:  true,
		compress:   false,
	}
	for _, option := range options {
		option(opts)
	}

	opts.maxAge = mathKit.Max(opts.maxAge, 0)
	opts.maxBackups = mathKit.Max(opts.maxBackups, 0)

	return opts
}

// NewLumberJackWriteCloser 可rotate（依据传参maxSize）的io.WriteCloser.
/*
PS:
(1) 如果 当前目标文件 被人为删了，会丢失部分输出直至再次调用Rotate().
(2) 可以同时配置 MaxBackups 和 MaxAge.

@param maxSize unit: byte
@param options 可选配置:
WithCompress()		[默认: 不压缩] 是否压缩被rotate的文件？
WithMaxBackups()	[默认: 保留所有旧日志] 旧日志保存的最多数量
WithMaxAge()		[默认: 保留所有旧日志] 旧日志保存的最长时间
WithLocalTime()		[默认: true] true: 使用本地时间; false: 使用UTC时间
*/
func NewLumberJackWriteCloser(filePath string, maxSize int64, options ...LumberJackOption) (*lumberjack.Roller, error) {
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

	opts := loadOptions(options...)
	if opts.maxAge > 0 && opts.maxAge < time.Minute {
		return nil, errorKit.Newf("maxAge(%s) is too small", opts.maxAge)
	}
	return lumberjack.NewRoller(filePath, maxSize, &lumberjack.Options{
		MaxAge:     opts.maxAge,
		MaxBackups: opts.maxBackups,
		LocalTime:  opts.localTime,
		Compress:   opts.compress,
	})
}
