package console

import (
	"os"

	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/richelieu042/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	mutex = new(gmutex.RWMutex)

	// defLevel 默认日志级别: DEBUG
	defLevel = zap.DebugLevel

	l       *zap.Logger
	sl      *zap.SugaredLogger
	innerL  *zap.Logger
	innerSL *zap.SugaredLogger
)

func init() {
	initialize()
}

// initialize
/*
调用此函数前须确保：已经获取了写锁 or 由init()调用
*/
func initialize() {
	encoder := zapKit.NewEncoder()
	ws := os.Stdout
	core := zapKit.NewCore(encoder, ws, defLevel)

	l = zapKit.NewLogger(core, zapKit.WithCallerSkip(0))
	sl = l.Sugar()
	innerL = zapKit.NewLogger(core, zapKit.WithCallerSkip(1))
	innerSL = innerL.Sugar()
}

// GetL 供外部使用（skip为0）
func GetL() *zap.Logger {
	/* 读锁 */
	mutex.RLock()
	defer mutex.RUnlock()

	return l
}

// GetSL 供外部使用（skip为0）
func GetSL() *zap.SugaredLogger {
	/* 读锁 */
	mutex.RLock()
	defer mutex.RUnlock()

	return sl
}

// getInnerL 供内部使用（skip为1）
func getInnerL() *zap.Logger {
	/* 读锁 */
	mutex.RLock()
	defer mutex.RUnlock()

	return innerL
}

// getInnerSL 供内部使用（skip为1）
func getInnerSL() *zap.SugaredLogger {
	/* 读锁 */
	mutex.RLock()
	defer mutex.RUnlock()

	return innerSL
}

func Sync() {
	/* 写锁 */
	mutex.LockFunc(func() {
		_ = l.Sync()
		_ = sl.Sync()
		_ = innerL.Sync()
		_ = innerSL.Sync()
	})
}

func SetLogLevel(level zapcore.Level) {
	/* 写锁 */
	mutex.LockFunc(func() {
		if level == defLevel {
			return
		}
		defLevel = level

		_ = l.Sync()
		_ = sl.Sync()
		_ = innerL.Sync()
		_ = innerSL.Sync()

		initialize()
	})
}
