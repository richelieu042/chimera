package console

import (
	"go.uber.org/zap"
)

func Debug(msg string, fields ...zap.Field) {
	getInnerL().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	getInnerL().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	getInnerL().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	getInnerL().Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	getInnerL().DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	getInnerL().Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	getInnerL().Fatal(msg, fields...)
}
