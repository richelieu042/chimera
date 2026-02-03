package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"go.uber.org/zap/zapcore"
)

// NewCore
/*
@param encoder		决定日志格式（不能为nil）
@param ws			决定日志写入位置，如文件、控制台、网络等（可以为nil: 默认输出到控制台）
@param levelEnabler	（1）决定哪些日志级别会被记录（不能为nil）
					（2）可以是多种:
						(a) zapcore.Level 类型（级别 >= 此值的才会输出）
							e.g.
							zapcore.DebugLevel
							zapcore.InfoLevel
							zapcore.WarnLevel
							zapcore.ErrorType
							zapcore.PanicLevel
							zapcore.DPanicLevel
							zapcore.FatalLevel
							zapcore.InvalidLevel
						(b) zap.LevelEnablerFunc 类型（更加地自定义）
							e.g.
							// 创建错误日志级别的核心
							errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
								return level >= zapcore.ErrorLevel
							})
@param initialFields 可以不传
*/
func NewCore(encoder zapcore.Encoder, writeSyncer zapcore.WriteSyncer, levelEnabler zapcore.LevelEnabler, initialFields ...zapcore.Field) zapcore.Core {
	if writeSyncer == nil {
		writeSyncer = LockedWriteSyncerStdout
	}

	core := zapcore.NewCore(encoder, writeSyncer, levelEnabler)
	if len(initialFields) > 0 {
		core = core.With(initialFields)
	}
	return core
}

func NewLazyWith(core zapcore.Core, fields []zapcore.Field) zapcore.Core {
	return zapcore.NewLazyWith(core, fields)
}

func NewIncreaseLevelCore(core zapcore.Core, levelEnabler zapcore.LevelEnabler) (zapcore.Core, error) {
	return zapcore.NewIncreaseLevelCore(core, levelEnabler)
}

// MultiCore
/*
@return != nil
*/
func MultiCore(cores ...zapcore.Core) zapcore.Core {
	cores = sliceKit.RemoveZeroValues(cores)

	return zapcore.NewTee(cores...)
}
