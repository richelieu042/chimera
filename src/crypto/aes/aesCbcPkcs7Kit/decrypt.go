package aesCbcPkcs7Kit

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu042/chimera/v3/src/crypto/hexKit"
)

// Decrypt AES/CBC/PKCS7 解密.
func Decrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(data))
	blockMode.CryptBlocks(origData, data)
	origData = pkcs7UnPadding(origData)
	return origData, nil
}

func DecryptFromBase64(base64Str string, key []byte, iv []byte) ([]byte, error) {
	data, err := base64Kit.DecodeString(base64Str, base64Kit.WithEncoding(base64.StdEncoding))
	if err != nil {
		return nil, errorKit.Wrapf(err, "Fail to decode as base64 string")
	}

	return Decrypt(data, key, iv)
}

func DecryptFromHex(hexStr string, key []byte, iv []byte) ([]byte, error) {
	data, err := hexKit.DecodeString(hexStr)
	if err != nil {
		return nil, errorKit.Wrapf(err, "Fail to decode as hex string")
	}

	return Decrypt(data, key, iv)
}
