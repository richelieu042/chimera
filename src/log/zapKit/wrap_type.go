package zapKit

import (
	"io"

	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"go.uber.org/zap"
)

type (
	WrappedLogger struct {
		*zap.Logger

		Writers []io.Writer
	}

	WrappedSugaredLogger struct {
		*zap.SugaredLogger

		Writers []io.Writer
	}
)

// Close 释放资源.
/*
!!!: 要注意并发问题，调用本方法后，不要再通过此logger输出了.
*/
func (l *WrappedLogger) Close() (err error) {
	if l == nil {
		return
	}

	_ = l.Sync()
	for _, w := range l.Writers {
		tmpErr := ioKit.TryToClose(w)
		if tmpErr != nil && err == nil {
			err = tmpErr
		}
	}
	return
}

// Close 释放资源.
/*
!!!: 要注意并发问题，调用本方法后，不要再通过此logger输出了.
*/
func (sl *WrappedSugaredLogger) Close() (err error) {
	if sl == nil {
		return
	}

	_ = sl.Sync()
	for _, w := range sl.Writers {
		tmpErr := ioKit.TryToClose(w)
		if tmpErr != nil && err == nil {
			err = tmpErr
		}
	}
	return
}
