package aesCbcPkcs7Kit

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu042/chimera/v3/src/crypto/hexKit"
)

// Encrypt AES/CBC/PKCS7 加密.
/*
PS: 返回值转换为字符串，可以考虑使用: base64.StdEncoding.EncodeToString().
*/
func Encrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	data = pkcs7Padding(data, blockSize)
	crypted := make([]byte, len(data))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(crypted, data)
	return crypted, nil
}

// EncryptToBase64 PS: 标准的base64.
func EncryptToBase64(data []byte, key []byte, iv []byte) (string, error) {
	data, err := Encrypt(data, key, iv)
	if err != nil {
		return "", err
	}

	return base64Kit.EncodeToString(data, base64Kit.WithEncoding(base64.StdEncoding)), nil
}

func EncryptToHex(data []byte, key []byte, iv []byte) (string, error) {
	data, err := Encrypt(data, key, iv)
	if err != nil {
		return "", err
	}

	return hexKit.EncodeToString(data), nil
}
