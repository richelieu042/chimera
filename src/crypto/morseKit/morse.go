package morseKit

import (
	"github.com/alwindoss/morse"
	"github.com/richelieu042/chimera/v3/src/core/ioKit"
)

// Encode
/*
Deprecated: 英文字母不区分大小写，且和在线编码解码不一致.

@param plainData 支持中文
*/
func Encode(plainData []byte) (cryptoData []byte, err error) {
	reader := ioKit.NewReader(plainData)
	hacker := morse.NewHacker()
	return hacker.Encode(reader)
}

// Decode
/*
Deprecated: 英文字母不区分大小写，且和在线编码解码不一致.
*/
func Decode(cryptoData []byte) (plainData []byte, err error) {
	reader := ioKit.NewReader(cryptoData)
	hacker := morse.NewHacker()
	return hacker.Decode(reader)
}
