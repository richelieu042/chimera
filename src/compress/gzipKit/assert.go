package gzipKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

func AssertValidLevel(level int) error {
	if !IsValidLevel(level) {
		return errorKit.NewfWithSkip(1, "[%s] invalid gzip compression level(%d)", funcKit.GetFuncName(1), level)
	}
	return nil
}
