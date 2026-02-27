package zapKit

import (
	"io"
	"os"

	"go.uber.org/zap/zapcore"
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
(0) 对于 os.Stderr 和 os.Stdout，不需要 zapcore.Lock;
(1) os.File 结构体实现了 zapcore.WriteSyncer 接口;
(2) zapcore.WriteSyncer 接口是 io.Writer 接口的子类.
*/
func NewLockedWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	if w == nil {
		return nil
	}

	switch w {
	case os.Stdout, os.Stderr:
		return zapcore.AddSync(w)
	default:
		ws := zapcore.AddSync(w)
		return zapcore.Lock(ws)
	}
}

// MultiWriteSyncer 类似于 io.MultiWriter.
func MultiWriteSyncer(ws ...zapcore.WriteSyncer) zapcore.WriteSyncer {
	if len(ws) == 0 {
		return nil
	}

	return zapcore.NewMultiWriteSyncer(ws...)
}
