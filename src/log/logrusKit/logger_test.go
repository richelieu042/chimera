package logrusKit

import (
	"testing"

	"github.com/richelieu042/chimera/v3/src/core/ioKit"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger(WithMsgPrefix("[TEST] "))
	logger.Info("Hello world.")
}

func TestNewLogger1(t *testing.T) {
	/* 输出到控制台 */
	logger := NewLogger()
	logger.Info("to defaultOutput")

	/* 输出到文件(rotatable) */
	path := "_test.log"
	writer, err := ioKit.NewLumberjackWriteCloser(path)
	if err != nil {
		panic(err)
	}
	logger = NewLogger(WithOutput(writer))
	logger.Info("rotatable file logger")
}

func TestNewFileLogger(t *testing.T) {
	/* 输出到文件(not rotatable) */
	logger, err := NewFileLogger("_test1.log")
	if err != nil {
		panic(err)
	}

	logger.Info("test")
	logger.Info("测试")
	logger.Info("~！@#￥%……&*（）——+`1234567890-=")
}
