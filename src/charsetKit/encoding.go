package charsetKit

import (
	"unicode/utf8"

	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

// IsSupported 是否支持 指定字符集 ？
var IsSupported func(charset string) bool = gcharset.Supported

// DetermineEncoding 编码推断.
/*
Deprecated: 不一定准（检测GBK文本有问题，推断出来的是 "windows-1252"）.
*/
var DetermineEncoding func(content []byte, contentType string) (e encoding.Encoding, name string, certain bool) = charset.DetermineEncoding

// Detect 检测文本的编码.
/*
Deprecated: 不一定准（检测GBK文本有问题，推断出来的是 "ISO-8859-1"）.
*/
func Detect(data []byte) (charset string, err error) {
	if err = interfaceKit.AssertNotNil(data, "data"); err != nil {
		return
	}

	var r *chardet.Result
	r, err = chardet.NewTextDetector().DetectBest(data)
	if err != nil {
		return
	}
	return r.Charset, nil
}

var IsUTF8 func(data []byte) bool = utf8.Valid

var IsUTF8String func(s string) bool = utf8.ValidString

// IsGBK
/*
PS: 在Go语言中，你可以通过检查字节序列是否落在 GBK 编码范围内来判断文本编码是否为 GBK。
*/
func IsGBK(data []byte) bool {
	if IsUTF8(data) {
		return false
	}

	i := 0
	for i < len(data) {
		if data[i] <= 0x7f {
			// 编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		}
		if data[i] >= 0x81 && data[i] <= 0xfe && data[i+1] >= 0x40 && data[i+1] <= 0xfe && data[i+1] != 0xf7 {
			// 大于127的使用双字节编码，落在gbk编码范围内的字符
			i += 2
			continue
		}
		return false
	}
	return true
}
