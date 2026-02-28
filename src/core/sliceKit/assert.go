package sliceKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

// AssertNotEmpty
/*
合法的传参:
(1) s != nil
(2) len(s) > 0
*/
func AssertNotEmpty[T any](s []T, name string) error {
	if len(s) == 0 {
		if s == nil {
			return errorKit.NewfWithSkip(1, "[%s] param(name: %s, type: %s) == nil",
				funcKit.GetFuncName(1), name, "slice")
		}
		return errorKit.NewfWithSkip(1, "[%s] param(name: %s, type: %s) is empty",
			funcKit.GetFuncName(1), name, "slice")
	}
	return nil
}
