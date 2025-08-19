package logrusKit

import (
	"os"

	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"github.com/sirupsen/logrus"
)

// NewTruncFileLogger
/*
@param filePath (1) 如果不存在，会自动创建;
				(2) 如果 已经存在 && 是个文件，会截断该文件的内容.
*/
func NewTruncFileLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	file, err := fileKit.Create(filePath)
	if err != nil {
		return nil, err
	}
	options = append(options, WithOutput(file))

	return NewLogger(options...), nil
}

// NewTruncFileAndStdoutLogger
/*
@param filePath (1) 如果不存在，会自动创建;
				(2) 如果 已经存在 && 是个文件，会截断该文件的内容.
*/
func NewTruncFileAndStdoutLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	f, err := fileKit.Create(filePath)
	if err != nil {
		return nil, err
	}
	output := ioKit.MultiWriter(f, os.Stdout)
	options = append(options, WithOutput(output))

	return NewLogger(options...), nil
}
