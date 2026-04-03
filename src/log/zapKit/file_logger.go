package zapKit

import (
	"os"

	"github.com/gogf/gf/v2/os/gfile"
	"go.uber.org/zap/zapcore"
)

// NewFileLogger
/*
PS:
（1）不再使用返回的 Logger 时，记得手动调用 Close() ！！！
（2）仅输出到文件；
（3）不要调用 fileKit，以免 import cycle.

@param filePath 日志文件路径
*/
func NewFileLogger(filePath string, prefix string, level zapcore.Level, loggerOptions ...LoggerOption) (*WrappedLogger, error) {
	/* 参考了：fileKit.CreateInAppendMode */
	// 创建父目录（如果不存在的话）
	parentDir := gfile.Dir(filePath)
	if !gfile.Exists(parentDir) {
		if err := gfile.Mkdir(parentDir); err != nil {
			return nil, err
		}
	}
	// 打开文件（此处不要关闭文件，已关闭就写不了日志了）
	f, err := gfile.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) // 如果文件有内容，会append
	if err != nil {
		return nil, err
	}

	//f, err := fileKit.CreateInAppendMode(filePath)
	//if err != nil {
	//	return nil, err
	//}

	enc := NewEncoder(WithEncoderMessagePrefix(prefix))
	core := NewCore(enc, zapcore.Lock(f), level)
	logger := NewLogger(core, loggerOptions...)

	return WrapLogger(logger, f), nil
}
