package reflectKit

import (
	"reflect"
	"unsafe"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
)

// GetNestedField 获取（层层嵌套的）字段
func GetNestedField(ptr interface{}, fieldNames ...string) (reflect.Value, error) {
	if fieldNames == nil {
		return reflect.Value{}, errorKit.Newf("fieldNames == nil")
	}

	v := reflect.ValueOf(ptr).Elem()
	for i, name := range fieldNames {
		if i > 0 {
			/*
				参考: 	https://www.codenong.com/50098624/
				以避免: 	panic: reflect: call of reflect.Value.FieldByName on ptr Value
			*/
			v = reflect.Indirect(v)
		}
		v = v.FieldByName(name)
	}
	return v, nil
}

// SetField 设置字段的值.
/*
PS: 支持设置 未导出字段 的值.

@param fieldName 		字段名（属性名），private或public的都可以
@param newFieldValue	字段的新值
*/
func SetField(ptr interface{}, fieldName string, newFieldValue interface{}) error {
	v := GetField(ptr, fieldName)
	// 获取字段的可寻址反射对象（可set）
	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()

	v1 := reflect.ValueOf(newFieldValue)
	if v.Kind() != v1.Kind() {
		return errorKit.Newf("expected kind %v, got kind: %v", v.Kind(), v1.Kind())
	}
	v.Set(v1)
	return nil
}
