package typeKit

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
)

// GetTypeString 获取对象的类型字符串.
/*
@param 要注意nil的情况!!!
*/
func GetTypeString[T any](obj T) string {
	if interfaceKit.IsNil(obj) {
		return "<nil>"
	}

	return fmt.Sprintf("%T", obj)
}
