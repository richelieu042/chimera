package ptrKit

import (
	"fmt"
	"reflect"

	"github.com/modern-go/reflect2"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// IsPointer 传参的类型是否为指针？
/*
PS: 传参的值无影响，即使其为nil.

@param obj 可以为nil

e.g.
	type bean struct {
	}
	var b *bean = nil
	fmt.Println(ptrKit.IsPointer(b)) 				// true（类型为指针，虽然值为nil）

	fmt.Println(ptrKit.IsPointer(interface{}(nil))) // false
*/
func IsPointer(obj interface{}) bool {
	// 参数obj为nil的话，变量v不为nil
	v := reflect.ValueOf(obj)
	return v.Kind() == reflect.Ptr
}

// IsPointer1
/**
 * 判断参数obj是否为指针（pointer）.
 * 参考：github.com/json-iterator/go => reflect.go => func (iter *Iterator) ReadVal(obj interface{})
 *
 * @param obj nil => false
 */
func IsPointer1(obj interface{}) bool {
	typ := reflect2.TypeOf(obj)
	return typ != nil && typ.Kind() == reflect.Ptr
}

// IsPointer2
/*
e.g.
(nil) => false
*/
func IsPointer2(obj interface{}) bool {
	str := fmt.Sprintf("%T", obj)
	return strKit.StartWith(str, "*")
}
