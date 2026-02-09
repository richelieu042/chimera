package logKit

import (
	"io"
	"log"
	"os"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// NewLogger
/*
   @param flag e.g. os.O_CREATE|os.O_WRONLY|os.O_APPEND
*/
func NewLogger(out io.Writer, prefix string, flag int) *log.Logger {
	return log.New(out, prefix, flag)
}

// NewStdoutLogger 输出到控制台（os.Stdout）.
func NewStdoutLogger(prefix string) *log.Logger {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	if strKit.IsNotEmpty(prefix) {
		flag |= log.Lmsgprefix
	}
	return NewLogger(os.Stdout, prefix, flag)
}

// NewFileLogger
/*
@param logPath 	(1) 文件不存在，会生成文件和父目录;
				(2) 文件存在，新的内容会 append.
@param prefix 	e.g. "[TEST] "
@param perm 	e.g. 0666 || 0644
*/
func NewFileLogger(filePath, prefix string, perm os.FileMode) (*log.Logger, error) {
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

	/* out */
	// 如果生成文件的话，其权限为0666（与os.Create()的权限一样）
	// os.OpenFile()的传参flag，可参考：https://studygolang.com/articles/22180
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, perm)
	if err != nil {
		return nil, err
	}
	// 此处不能关闭writer，否则日志内容将写不进去
	//defer f.Close()

	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	if strKit.IsNotEmpty(prefix) {
		flag |= log.Lmsgprefix
	}
	return NewLogger(f, prefix, flag), nil
}
