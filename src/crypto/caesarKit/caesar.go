package caesarKit

import (
	"encoding/base64"

	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
)

// Encrypt 凯撒密码加密.
/*
PS: 仅加密大小写的英文字母，不会动中文.

@param shift 推荐值: (1) [1, 25]
					(2) 不推荐使用 26*n，因为这样加密了个寂寞
*/
func Encrypt(plainText string, shift int) (cipherText string) {
	cipher := NewCaesarCipher(shift)
	cipherText = cipher.Encrypt(plainText)
	return
}

func Decrypt(cipherText string, shift int) (plainText string) {
	cipher := NewCaesarCipher(shift)
	plainText = cipher.Decrypt(cipherText)
	return
}

// EncryptWithRawURLBase64 凯撒密码加密 + base64 编码.
/*
PS:
(1) base64采用了 base64.RawURLEncoding;
(2) 传参 plainText 可以有中文;
(3) 返回值可以放到 url 中（但不推荐）.
*/
func EncryptWithRawURLBase64(plainText string, shift int) (cipherText string) {
	// base64 编码
	base64Str := base64Kit.EncodeStringToString(plainText, base64Kit.WithEncoding(base64.RawURLEncoding))
	// 凯撒密码 加密
	cipherText = Encrypt(base64Str, shift)
	return
}

func DecryptWithRawURLBase64(cipherText string, shift int) (plainText string, err error) {
	// 凯撒密码 解密
	base64Str := Decrypt(cipherText, shift)
	// base64 解码
	plainText, err = base64Kit.DecodeStringToString(base64Str, base64Kit.WithEncoding(base64.RawURLEncoding))
	return
}
