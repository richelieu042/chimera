package intKit

import (
	"strconv"

	"github.com/spf13/cast"
)

var (
	ToInt func(i interface{}) int = cast.ToInt

	// ToIntE
	/*
		PS:
		(1) 2进制的前缀:	0b
		(2) 8进制的前缀:	0
		(3) 10进制的前缀:	无
		(4) 16进制的前缀:	0x || 0X（字母x不区分大小写，即0x与0X等价）

		e.g.
			fmt.Println(strconv.Atoi("08"))  	// 8 <nil>
			fmt.Println(strconv.Atoi("0xg")) 	// 0 strconv.Atoi: parsing "0xg": invalid syntax

			// 0开头是"8进制数"，8进制数不应该出现"8"，所以转换失败
			fmt.Println(cast.ToIntE("08")) 		// 0 unable to cast "08" of type string to int64
			// 0x开头是"16进制数"，16进制数不应该出现"g"，所以转换失败
			fmt.Println(cast.ToIntE("0xg")) 	// 0 unable to cast "0xg" of type string to int64

		e.g.1
		   (nil) 	=> 0 <nil>
		   (false)	=> 0 <nil>
		   (true)	=> 1 <nil>
		   ("")		=> 0 unable to cast "" of type string to int64
	*/
	ToIntE func(i interface{}) (int, error) = cast.ToIntE

	ToInt8 func(i interface{}) int8 = cast.ToInt8

	ToInt8E func(i interface{}) (int8, error) = cast.ToInt8E

	ToInt16 func(i interface{}) int16 = cast.ToInt16

	ToInt16E func(i interface{}) (int16, error) = cast.ToInt16E

	ToInt32 func(i interface{}) int32 = cast.ToInt32

	ToInt32E func(i interface{}) (int32, error) = cast.ToInt32E

	ToInt64 func(i interface{}) int64 = cast.ToInt64

	ToInt64E func(i interface{}) (int64, error) = cast.ToInt64E

	// StringToInt 类型转换: string => int
	/*
		PS:
		(1) 传参s 必须是10进制的数字字符串（别的进制会返回error）.
		(2) strconv.Atoi <=> strconv.ParseInt(s, 10, 0)
	*/
	StringToInt func(s string) (int, error) = strconv.Atoi

	// StringToIntWithBase
	/*
		@param base		传参s的进制数（2 || 8 || 10 || 16）
		@param bitSize 	返回值的位数（0 || 8 || 16 || 32 || 64; 0: 使用当前系统的位数）
	*/
	StringToIntWithBase func(s string, base int, bitSize int) (i int64, err error) = strconv.ParseInt
)

// StringToIntWithDefault 类型转换: string => int

func StringToIntWithDefault(str string, def int) int {
	i, err := StringToInt(str)
	if err != nil {
		return def
	}
	return i
}
