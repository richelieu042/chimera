// Package httpKit
/*
请求头 || 响应头
key: 不区分大小写
*/
package httpKit

import (
	"bytes"
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/mapKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// HeaderToMap http.Header（即 map[string][]string） => map[string]interface{}
func HeaderToMap(header http.Header) map[string]interface{} {
	return mapKit.MapValues(header, func(value []string, key string) interface{} {
		return value
	})
}

// HeaderToString Converts http header to a string.
func HeaderToString(header http.Header) (string, error) {
	if header == nil {
		return "", nil
	}

	buf := new(bytes.Buffer)
	if err := header.Write(buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// SetHeader
/*
e.g.
	header := make(map[string][]string)

	httpKit.AddHeader(header, "k", "0")
	fmt.Println(header) 				// map[K:[0]]
	httpKit.AddHeader(header, "k", "1")
	fmt.Println(header) 				// map[K:[0 1]]

	httpKit.SetHeader(header, "k", "2")
	fmt.Println(header) 				// map[K:[2]]
	httpKit.SetHeader(header, "k", "")
	fmt.Println(header) 				// map[]
*/
func SetHeader(header http.Header, key, value string) {
	if value == "" {
		DelHeader(header, key)
		return
	}
	header.Set(key, value)
}

// SetHeaderIfMissing
/*
PS: 区分大小写.
*/
func SetHeaderIfMissing(header http.Header, key, value string) {
	if !HeaderContainsValue(header, key, value) {
		SetHeader(header, key, value)
	}
}

// SetHeaderIfMissingIgnoreCase
/*
PS: 不区分大小写.
*/
func SetHeaderIfMissingIgnoreCase(header http.Header, key, value string) {
	if !HeaderContainsValueIgnoreCase(header, key, value) {
		SetHeader(header, key, value)
	}
}

// GetHeader
/*
PS:
(1) 不存在与 key 对应值的话，将返回 "";
(2) 存在多个值的话（[]string && len > 1），返回第一个值.
*/
func GetHeader(header http.Header, key string) string {
	return header.Get(key)
}

// GetHeaderValues
/*
@param key 不存在对应值的话，将返回 nil
*/
func GetHeaderValues(header http.Header, key string) []string {
	return header.Values(key)
}

func DelHeader(header http.Header, key string) {
	header.Del(key)
}

// AddHeader
/*
PS: key对应的 切片 中，在 最后面 添加value.

e.g.
	header := make(map[string][]string)
	AddHeader(header, "k", "0")
	fmt.Println(header) 		// map[K:[0]]
	AddHeader(header, "k", "1")
	fmt.Println(header) 		// map[K:[0 1]]
	AddHeader(header, "k", "1")
	fmt.Println(header) 		// map[K:[0 1 1]]
*/
func AddHeader(header http.Header, key, value string) {
	header.Add(key, value)
}

// AddHeaderIfMissing
/*
PS: 区分大小写.
*/
func AddHeaderIfMissing(header http.Header, key, value string) {
	if !HeaderContainsValue(header, key, value) {
		AddHeader(header, key, value)
	}
}

// AddHeaderIfMissingIgnoreCase
/*
PS: 不区分大小写.
*/
func AddHeaderIfMissingIgnoreCase(header http.Header, key, value string) {
	if !HeaderContainsValueIgnoreCase(header, key, value) {
		AddHeader(header, key, value)
	}
}

// HeaderContainsValue Header中，指定key对应的value切片是否包含指定value？
/*
PS: 区分大小写.
*/
func HeaderContainsValue(header http.Header, key, value string) bool {
	values := header.Values(key)
	return sliceKit.Contains(values, value)
}

// HeaderContainsValueIgnoreCase Header中，指定key对应的value切片是否包含指定value？\
/*
PS: 不区分大小写.
*/
func HeaderContainsValueIgnoreCase(header http.Header, key, value string) bool {
	values := header.Values(key)
	return sliceKit.ContainsBy(values, func(item string) bool {
		return strKit.EqualsIgnoreCase(item, value)
	})
}
