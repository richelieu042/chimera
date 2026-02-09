package aesKit

import (
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
)

// EncryptCbcPKCS5PaddingToString 明文([]byte) => 密文(string)
/*
@param key 	密钥（must be 16/24/32 bit length）
@param iv	偏移量，	(1) 默认: []byte("I Love Go Frame!")
					(2) 建议长度为16（由于源码中的 aes.BlockSize）
*/
func EncryptCbcPKCS5PaddingToString(plainData []byte, key []byte, iv ...[]byte) (string, error) {
	cipherData, err := EncryptCbcPKCS5Padding(plainData, key, iv...)
	if err != nil {
		return "", err
	}
	// 转换成base64字符串
	return base64Kit.EncodeToString(cipherData), nil
}

// DecryptCbcPKCS5PaddingFromString 密文(string) => 明文([]byte)
func DecryptCbcPKCS5PaddingFromString(base64Str string, key []byte, iv ...[]byte) ([]byte, error) {
	cipherData, err := base64Kit.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	return DecryptCbcPKCS5Padding(cipherData, key, iv...)
}

// EncryptCbcPKCS5Padding
/*
加密模式: CBC
填充方式: PKCS5Padding
*/
var EncryptCbcPKCS5Padding func(plainData []byte, key []byte, iv ...[]byte) ([]byte, error) = gaes.EncryptCBC

// DecryptCbcPKCS5Padding
/*
加密模式: CBC
填充方式: PKCS5Padding
*/
var DecryptCbcPKCS5Padding func(cipherData []byte, key []byte, iv ...[]byte) ([]byte, error) = gaes.DecryptCBC
