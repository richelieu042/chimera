package ptrKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

// AssertNotNilAndIsPointer
/*
合法的传参:
(1) ptr != nil（值非nil）
(2) ptr的类型为"指针"
*/
func AssertNotNilAndIsPointer(ptr interface{}) error {
	if ptr == nil {
		return errKit.NewfWithDepth(1, "[%s] ptr == nil", funcKit.GetFuncName(1))
	}
	if !IsPointer(ptr) {
		return errKit.NewfWithDepth(1, "[%s] ptr(type: %T) isn't a pointer", funcKit.GetFuncName(1), ptr)
	}
	return nil
}
