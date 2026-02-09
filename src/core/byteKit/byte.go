// Package byteKit
/*
byte <=> uint8
*/
package byteKit

import (
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/spf13/cast"
)

func ToByte(obj interface{}) byte {
	return ToByteWithDefault(obj, 0)
}

func ToByteWithDefault(obj interface{}, def byte) byte {
	b, _ := ToByteWithDefaultE(obj, def)
	return b
}

func ToByteWithDefaultE(obj interface{}, def byte) (byte, error) {
	switch obj.(type) {
	case string:
		return stringToByte(obj.(string), def)
	default:
		if obj == nil {
			return def, nil
		}
		rst, err := cast.ToUint8E(obj)
		if err != nil {
			return def, err
		}
		return rst, nil
	}
}

func stringToByte(str string, def byte) (byte, error) {
	str = strKit.TrimSpace(str)

	if strKit.IsEmpty(str) {
		return def, nil
	}
	rst, err := cast.ToUint8E(str)
	if err != nil {
		return def, err
	}
	return rst, nil
}
