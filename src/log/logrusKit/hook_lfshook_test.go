package logrusKit

import (
	"os"
	"testing"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// TestNewLfsHook
/*
输出到文件: 	所有级别
输出到控制台: ERROR、PANIC、FATAL
*/
func TestNewLfsHook(t *testing.T) {
	logger, err := NewTruncFileLogger("_test_lfshook.log")
	if err != nil {
		panic(err)
	}

	hook := NewLfsHook(lfshook.WriterMap{
		logrus.WarnLevel:  os.Stdout,
		logrus.ErrorLevel: os.Stdout,
		logrus.PanicLevel: os.Stdout,
		logrus.FatalLevel: os.Stdout,
	}, logger.Formatter)
	logger.AddHook(hook)

	logger.Debugf("Debug %d", 0)
	logger.Infof("Info %d", 1)
	logger.Warnf("Warn %d", 2)
	logger.Errorf("Error %d", 3)

	//logger.Panicf("Panic %d", 4)
	//logger.Fatalf("Fatal %d", 5)
}
