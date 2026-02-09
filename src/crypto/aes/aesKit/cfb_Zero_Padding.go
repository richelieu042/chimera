package aesKit

import (
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
)

// EncryptCfbZeroPaddingToString 明文([]byte) => 密文(string)
/*
@param key 	密钥（must be 16/24/32 bit length）
@param iv	偏移量，	(1) 默认: []byte("I Love Go Frame!")
					(2) 建议长度为16（由于源码中的 aes.BlockSize）
*/
func EncryptCfbZeroPaddingToString(plainData []byte, key []byte, padding *int, iv ...[]byte) (string, error) {
	cipherData, err := EncryptCfbZeroPadding(plainData, key, padding, iv...)
	if err != nil {
		return "", err
	}
	// 转换成base64字符串
	return base64Kit.EncodeToString(cipherData), nil
}

// DecryptCfbZeroPaddingFromString 密文(string) => 明文([]byte)
func DecryptCfbZeroPaddingFromString(base64Str string, key []byte, unPadding int, iv ...[]byte) ([]byte, error) {
	cipherData, err := base64Kit.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	return DecryptCfbZeroPadding(cipherData, key, unPadding, iv...)
}

// EncryptCfbZeroPadding
/*
加密模式: CFB
填充方式: ZeroPadding
*/
var EncryptCfbZeroPadding func(plainData []byte, key []byte, padding *int, iv ...[]byte) ([]byte, error) = gaes.EncryptCFB

// DecryptCfbZeroPadding
/*
加密模式: CFB
填充方式: ZeroPadding
*/
var DecryptCfbZeroPadding func(cipherData []byte, key []byte, unPadding int, iv ...[]byte) ([]byte, error) = gaes.DecryptCFB
