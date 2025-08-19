package zapKit

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewEncoder(t *testing.T) {
	enc := NewEncoder()
	core := zapcore.NewCore(enc, nil, zapcore.DebugLevel)
	logger := zap.New(core)

	logger.Debug("This is a debug message", zap.String("key", "value"))
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))
}
