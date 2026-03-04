package console

import (
	"testing"

	"go.uber.org/zap"
)

func TestSetLogLevel(t *testing.T) {
	Debug("DEBUG")
	Info("INFO")
	Warn("WARN")
	Error("ERROR")

	SetLogLevel(zap.WarnLevel)

	Debug("DEBUG", zap.Int("int", 1))
	Info("INFO", zap.Int("int", 1))
	Warn("WARN", zap.Int("int", 1))
	Error("ERROR", zap.Int("int", 1))
}
