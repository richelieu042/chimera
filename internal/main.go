package main

import (
	"github.com/richelieu042/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
)

func main() {
	logger := zapKit.NewSimpleConsoleLogger()

	logger.Info("hello world", zap.Bool("flag", true), zap.Float64("f", 3.1415926))
	logger.Sugar().Info("hello world", zap.Bool("flag", true), zap.Float64("f", 3.1415926))
}
