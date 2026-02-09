package validateKit

import (
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/richelieu042/chimera/v3/src/reflectKit"
)

const (
	// MaxPort 65535 == 0xFFFF
	MaxPort = 0xFFFF
)

var bakedInValidators = map[string]validator.Func{
	"port":        isPort,
	"file_if":     isFileIf,
	"file_unless": isFileUnless,
}

func isPort(fl validator.FieldLevel) bool {
	field := fl.Field()
	return isValidPort(field)
}

func registerBakedInValidation(v *validator.Validate) (err error) {
	for tag, fn := range bakedInValidators {
		err = v.RegisterValidation(tag, fn)
		if err != nil {
			return
		}
	}
	return
}

// isValidPort
/*
参考:
(1) Java，hutool中的NetUtil.isValidPort()
(2) Linux端口分配: https://blog.csdn.net/zh2508/article/details/104888743

0 			不使用
1–1023 		系统保留,只能由root用户使用
1024—4999 	由客户端程序自由分配
5000—65535 	由服务器端程序自由分配（65535 = 2 ^ 16 - 1）

@param obj 	(1) 支持的类型: reflect.Value、int、uint、string...
			(2) 可以为nil
*/
func isValidPort(obj interface{}) bool {
	if obj == nil {
		return false
	}

	if v, ok := obj.(reflect.Value); ok {
		return isReflectValueValidPort(v)
	}
	return isReflectValueValidPort(reflectKit.ValueOf(obj))
}

func isReflectValueValidPort(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		i, err := strconv.Atoi(v.String())
		if err != nil {
			return false
		}
		return i > 0 && i <= MaxPort
	default:
		if v.CanInt() {
			i := v.Int()
			return i > 0 && i <= MaxPort
		} else if v.CanUint() {
			i := v.Uint()
			return i > 0 && i <= MaxPort
		}
		return false
	}
}
