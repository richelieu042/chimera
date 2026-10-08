package interfaceKit

import "reflect"

// IsNil Deprecated: 反射有性能问题，应尽可能避免使用此方法.
/*
go里面的类型包含 type 和 value 的，一般对于业务我们更在意value是不是空.

e.g.
	var src interface{} = nil
	var src1 []string = nil
	var src2 map[string]interface{} = nil
	type bean struct {
	}
	var src3 *bean = nil

	fmt.Println(interfaceKit.IsNil(src))  // true
	fmt.Println(interfaceKit.IsNil(src1)) // true
	fmt.Println(interfaceKit.IsNil(src2)) // true
	fmt.Println(interfaceKit.IsNil(src3)) // true
*/
func IsNil(obj any) bool {
	if obj == nil {
		return true
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return value.IsNil()
	default:
		return false
	}
}
