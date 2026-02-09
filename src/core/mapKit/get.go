package mapKit

import (
	"github.com/richelieu042/chimera/v3/src/core/boolKit"
	"github.com/richelieu042/chimera/v3/src/core/intKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// Obtain
/*
@return (1) 第1个返回值: 如果key存在于m中，为key对应的值；否则为V类型的零值.
		(2) 第2个返回值: key是否存在于m中.

e.g.
	m := map[string]interface{}{
		"a": 0,
		"b": 1,
	}
	fmt.Println(mapKit.Obtain(m, "a")) // 0 true
	fmt.Println(mapKit.Obtain(m, "c")) // <nil> false
*/
func Obtain[K comparable, V any](m map[K]V, key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}

// Get
/*
@param m 	可以为nil
@param key	如果不存在对应的值，将返回值类型的零值
*/
func Get[K comparable, V any](m map[K]V, key K) V {
	return m[key]
}

// GetString
/*
@param m 可以为nil
*/
func GetString[K comparable, V any](m map[K]V, key K) (string, error) {
	value := Get(m, key)
	return strKit.ToStringE(value)
}

// GetInt
/*
@param m 可以为nil
*/
func GetInt[K comparable, V any](m map[K]V, key K) (int, error) {
	value := Get(m, key)
	return intKit.ToIntE(value)
}

// GetInt32
/*
@param m 可以为nil
*/
func GetInt32[K comparable, V any](m map[K]V, key K) (int32, error) {
	value := Get(m, key)
	return intKit.ToInt32E(value)
}

// GetInt64
/*
@param m 可以为nil
*/
func GetInt64[K comparable, V any](m map[K]V, key K) (int64, error) {
	value := Get(m, key)
	return intKit.ToInt64E(value)
}

// GetBool
/*
@param m 可以为nil
*/
func GetBool[K comparable, V any](m map[K]V, key K) (bool, error) {
	value := Get(m, key)
	return boolKit.ToBoolE(value)
}
