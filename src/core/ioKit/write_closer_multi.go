package ioKit

import (
	"errors"
	"io"
	"sync"

	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
)

var (
	// closedError multiWriterCloser实例已经被关闭的情况下
	closedError = errors.New("multiWriterCloser has already been closed")
)

type multiWriterCloser struct {
	lock         *sync.Mutex
	writeClosers []io.WriteCloser
	closed       bool
}

func (multi *multiWriterCloser) Write(p []byte) (n int, err error) {
	multi.lock.Lock()
	defer multi.lock.Unlock()

	if multi.closed {
		return 0, closedError
	}

	for _, w := range multi.writeClosers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func (multi *multiWriterCloser) WriteString(s string) (n int, err error) {
	multi.lock.Lock()
	defer multi.lock.Unlock()

	if multi.closed {
		return 0, closedError
	}

	var p []byte // lazily initialized if/when needed
	for _, w := range multi.writeClosers {
		if sw, ok := w.(io.StringWriter); ok {
			n, err = sw.WriteString(s)
		} else {
			if p == nil {
				p = []byte(s)
			}
			n, err = w.Write(p)
		}
		if err != nil {
			return
		}
		if n != len(s) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(s), nil
}

func (multi *multiWriterCloser) Close() (err error) {
	multi.lock.Lock()
	defer multi.lock.Unlock()

	if multi.closed {
		return closedError
	}

	defer func() {
		multi.writeClosers = nil
		multi.closed = true
	}()

	// 修改元素类型(io.WriteCloser => io.Closer)
	closers := sliceKit.ConvertElementType(multi.writeClosers, func(item io.WriteCloser, _ int) io.Closer {
		return item
	})
	return Close(closers...)
}

// MultiWriteCloser
/*
PS: 参考了 io.MultiWriter().
*/
func MultiWriteCloser(writeClosers ...io.WriteCloser) io.WriteCloser {
	all := make([]io.WriteCloser, 0, len(writeClosers))

	for _, writeCloser := range writeClosers {
		if multi, ok := writeCloser.(*multiWriterCloser); ok {
			all = append(all, multi.writeClosers...)
		} else {
			all = append(all, writeCloser)
		}
	}
	return &multiWriterCloser{
		lock:         new(sync.Mutex),
		writeClosers: all,
	}
}
