package mapKit

import (
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

func AssertNotEmpty[K comparable, V any](m map[K]V, name string) error {
	if len(m) == 0 {
		if m == nil {
			return errorKit.NewfWithSkip(1, "[%s] param(name: %s, type: %s) == nil",
				funcKit.GetFuncName(1), name, "map")
		}
		return errorKit.NewfWithSkip(1, "[%s] param(name: %s, type: %s) is empty",
			funcKit.GetFuncName(1), name, "map")
	}
	return nil
}
