package ioKit

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

var (
	// NewReader
	/*
		PS: bytes.Reader 结构体实现了 io.Reader 接口.
	*/
	NewReader func(b []byte) *bytes.Reader = bytes.NewReader

	// NewReaderFromString
	/*
		PS: strings.Reader 结构体实现了 io.Reader 接口.
	*/
	NewReaderFromString func(s string) *strings.Reader = strings.NewReader
)

// NewReaderFromPath
/*
!!!: 要在外部手动调用 *os.File 的Close方法.

PS: os.File 结构体 实现了 io.Reader 接口.

@param path 文件（或目录）的路径
*/
func NewReaderFromPath(path string) (*os.File, error) {
	if err := fileKit.AssertExist(path); err != nil {
		return nil, err
	}
	return os.Open(path)
}

// NewBufferedReader 带缓冲的Reader.
/*
PS: bufio.Reader 结构体 实现了 io.Reader 接口.
*/
func NewBufferedReader(reader io.Reader, bufSizeArgs ...int) *bufio.Reader {
	if bufferedReader, ok := reader.(*bufio.Reader); ok {
		// 已经带缓冲了
		return bufferedReader
	}

	if len(bufSizeArgs) == 0 {
		// 默认缓冲大小: 4096
		return bufio.NewReader(reader)
	}
	return bufio.NewReaderSize(reader, bufSizeArgs[0])
}
