package zapKit

import (
	"strings"

	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// StringToLevel
/*
PS: 默认 DEBUG 级别.
*/
func StringToLevel(str string) (zapcore.Level, error) {
	if strKit.IsBlank(str) {
		return zapcore.DebugLevel, nil
	}

	str = strings.ToLower(str)
	return zapcore.ParseLevel(str)
}

// CanLoggerPrintSpecifiedLevel logger能否打印 指定级别lv 的日志？
func CanLoggerPrintSpecifiedLevel(logger *zap.Logger, lv zapcore.Level) bool {
	if logger == nil {
		return false
	}

	// 指定级别lv 大于等于 logger的级别，对应日志就会输出
	return lv >= logger.Level()
}
