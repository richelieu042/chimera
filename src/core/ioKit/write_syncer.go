package ioKit

import (
	"io"
	"os"

	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"go.uber.org/zap/zapcore"
)

var (
	// LockedWriteSyncerStdout （加锁的）标准输出.
	LockedWriteSyncerStdout = zapcore.Lock(os.Stdout)

	// LockedWriteSyncerStderr （加锁的）标准错误输出.
	LockedWriteSyncerStderr = zapcore.Lock(os.Stderr)
)

// NewWriteSyncer io.Writer => （并发不安全的）zapcore.WriteSyncer
/*
PS:
(1) os.File 结构体实现了 zapcore.WriteSyncer 接口;
(2) zapcore.WriteSyncer 接口是 io.Writer 接口的子类.
*/
func NewWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	if w == nil {
		return nil
	}

	return zapcore.AddSync(w)
}

// NewLockedWriteSyncer io.Writer => （并发安全的）zapcore.WriteSyncer
/*
PS:
(1) os.File 结构体实现了 zapcore.WriteSyncer 接口;
(2) zapcore.WriteSyncer 接口是 io.Writer 接口的子类.
*/
func NewLockedWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	if w == nil {
		return nil
	}

	switch w {
	case os.Stdout:
		return LockedWriteSyncerStdout
	case os.Stderr:
		return LockedWriteSyncerStderr
	default:
		return zapcore.Lock(zapcore.AddSync(w))
	}
}

// MultiWriteSyncer 类似于 io.MultiWriter.
func MultiWriteSyncer(ws ...zapcore.WriteSyncer) zapcore.WriteSyncer {
	ws = sliceKit.RemoveZeroValues(ws)
	if len(ws) == 0 {
		return nil
	}

	return zapcore.NewMultiWriteSyncer(ws...)
}
