package zapKit

import (
	"github.com/gogf/gf/v2/os/gmutex"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defMutex = new(gmutex.RWMutex)

	// defLevel 默认日志级别: DEBUG
	defLevel = zap.DebugLevel

	l       *zap.Logger
	sl      *zap.SugaredLogger
	innerL  *zap.Logger
	innerSL *zap.SugaredLogger
)

func init() {
	initializeLoggers()
}

func initializeLoggers() {
	encoder := NewEncoder()
	ws := LockedStdout
	core := NewCore(encoder, ws, defLevel)

	l = NewLogger(core, WithCallerSkip(0))
	sl = l.Sugar()
	innerL = NewLogger(core, WithCallerSkip(1))
	innerSL = innerL.Sugar()
}

// SetDefaultLevel PS: 默认日志级别为 DEBUG .
func SetDefaultLevel(level zapcore.Level) {
	/* 写锁 */
	defMutex.LockFunc(func() {
		if level == defLevel {
			return
		}
		defLevel = level

		initializeLoggers()
	})
}

func L() *zap.Logger {
	/* 读锁 */
	defMutex.RLock()
	defer defMutex.RUnlock()

	return l
}

func SL() *zap.SugaredLogger {
	/* 读锁 */
	defMutex.RLock()
	defer defMutex.RUnlock()

	return sl
}

func getInnerL() *zap.Logger {
	/* 读锁 */
	defMutex.RLock()
	defer defMutex.RUnlock()

	return innerL
}

func getInnerSL() *zap.SugaredLogger {
	/* 读锁 */
	defMutex.RLock()
	defer defMutex.RUnlock()

	return innerSL
}

func Sync() {
	/* 写锁 */
	defMutex.LockFunc(func() {
		_ = l.Sync()
		_ = sl.Sync()
		_ = innerL.Sync()
		_ = innerSL.Sync()
	})
}

func Debug(msg string, fields ...zap.Field) {
	getInnerL().Debug(msg, fields...)
}

// Info
/*
@param fields 输出循序与 传参fields 顺序一致（并不会按字母排序）
*/
func Info(msg string, fields ...zap.Field) {
	getInnerL().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	getInnerL().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	getInnerL().Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	getInnerL().Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	getInnerL().DPanic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	getInnerL().Fatal(msg, fields...)
}
