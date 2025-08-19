package zapKit

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewStdLogger *zap.Logger => *log.Logger
func NewStdLogger(l *zap.Logger) *log.Logger {
	return zap.NewStdLog(l)
}

// NewStdLoggerWithLevel *zap.Logger => *log.Logger
func NewStdLoggerWithLevel(l *zap.Logger, level zapcore.Level) (*log.Logger, error) {
	return zap.NewStdLogAt(l, level)
}

// RedirectStdLog 重定向 log标准库 的输出
func RedirectStdLog(l *zap.Logger) func() {
	return zap.RedirectStdLog(l)
}

func RedirectStdLogWithLevel(l *zap.Logger, level zapcore.Level) (func(), error) {
	return zap.RedirectStdLogAt(l, level)
}
