package interfaceKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

func AssertNotNil(obj interface{}, name string) error {
	if obj == nil {
		return errorKit.NewfWithSkip(1, "[%s] param(name: %s) == nil", funcKit.GetFuncName(1), name)
	}
	return nil
}
