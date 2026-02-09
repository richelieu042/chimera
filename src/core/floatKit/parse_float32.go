package floatKit

import (
	"strconv"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// ParseStringToFloat32 类型转换: string => float32
/*
e.g.
("3.141592653589793") 	=> 3.1415927, nil
("") 					=> 0, strconv.ParseFloat: parsing "": invalid syntax
*/
func ParseStringToFloat32(str string) (float32, error) {
	str = strKit.RemoveSpace(str)

	tmp, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(tmp), nil
}

func ParseStringToFloat32WithDefault(str string, def float32) float32 {
	rst, err := ParseStringToFloat32(str)
	if err != nil {
		return def
	}
	return rst
}
