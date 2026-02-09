package zapKit

import (
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"go.uber.org/zap/zapcore"
)

// NewFileLogger
/*
PS: 仅输出到文件.

@param filePath 日志文件路径
*/
func NewFileLogger(filePath string, prefix string, level zapcore.Level, loggerOptions ...LoggerOption) (*WrappedLogger, error) {
	f, err := fileKit.CreateInAppendMode(filePath)
	if err != nil {
		return nil, err
	}

	enc := NewEncoder(WithEncoderMessagePrefix(prefix))
	core := NewCore(enc, zapcore.Lock(f), level)
	logger := NewLogger(core, loggerOptions...)

	return WrapLogger(logger, f), nil
}
