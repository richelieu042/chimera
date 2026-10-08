package strKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

func AssertNotEmpty(str string, paramName string) error {
	if IsEmpty(str) {
		return errKit.NewfWithDepth(1, "[%s] param(name: %s) is empty",
			funcKit.GetFuncName(1), paramName)
	}
	return nil
}

func AssertNotBlank(str string, paramName string) error {
	if IsBlank(str) {
		return errKit.NewfWithDepth(1, "[%s] param(name: %s) is blank",
			funcKit.GetFuncName(1), paramName)
	}
	return nil
}
