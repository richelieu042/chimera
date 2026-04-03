package zapKit

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestNewFileLogger(t *testing.T) {
	path := "_test_file_logger.txt"

	l, err := NewFileLogger(path, "", zapcore.DebugLevel)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	l.Debug("Debug")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")
}
