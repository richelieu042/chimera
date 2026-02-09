package floatKit

import (
	"strconv"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// ParseStringToFloat64 类型转换: string => float64
/*
e.g.
("3.141592653589793") 	=> (3.141592653589793, nil)
("") 					=> 0, strconv.ParseFloat: parsing "": invalid syntax
*/
func ParseStringToFloat64(str string) (float64, error) {
	str = strKit.RemoveSpace(str)

	return strconv.ParseFloat(str, 64)
}

func ParseStringToFloat64WithDefault(str string, def float64) float64 {
	rst, err := ParseStringToFloat64(str)
	if err != nil {
		return def
	}
	return rst
}
