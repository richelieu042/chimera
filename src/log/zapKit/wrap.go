package zapKit

import (
	"io"

	"go.uber.org/zap"
)

// WrapLogger
/*
PS:
(1) 适用场景: 使用完后，需要释放对应资源（关闭输出）.
(2) 不要调用返回值的 Sugar()!!!
*/
func WrapLogger(logger *zap.Logger, writers ...io.Writer) *WrappedLogger {
	return &WrappedLogger{
		Logger:  logger,
		Writers: writers,
	}
}

// WrapSugarLogger
/*
PS:
(1) 适用场景: 使用完后，需要释放对应资源（关闭输出）.
(2) 不要调用返回值的 Desugar()!!!
*/
func WrapSugarLogger(sugaredLogger *zap.SugaredLogger, writers ...io.Writer) *WrappedSugaredLogger {
	return &WrappedSugaredLogger{
		SugaredLogger: sugaredLogger,
		Writers:       writers,
	}
}
