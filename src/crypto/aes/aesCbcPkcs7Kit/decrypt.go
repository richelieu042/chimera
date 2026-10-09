package aesCbcPkcs7Kit

import (
	"crypto/aes"
	"encoding/base64"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu042/chimera/v3/src/crypto/hexKit"
)

// Decrypt AES/CBC/PKCS7 解密.
func Decrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	if len(iv) != aes.BlockSize {
		return nil, errKit.Newf("invalid IV length(%d), want %d", len(iv), aes.BlockSize)
	}

	return gaes.DecryptCBC(data, key, iv)
}

func DecryptFromBase64(base64Str string, key []byte, iv []byte) ([]byte, error) {
	data, err := base64Kit.DecodeString(base64Str, base64Kit.WithEncoding(base64.StdEncoding))
	if err != nil {
		return nil, errKit.Wrapf(err, "Fail to decode as base64 string")
	}

	return Decrypt(data, key, iv)
}

func DecryptFromHex(hexStr string, key []byte, iv []byte) ([]byte, error) {
	data, err := hexKit.DecodeString(hexStr)
	if err != nil {
		return nil, errKit.Wrapf(err, "Fail to decode as hex string")
	}

	return Decrypt(data, key, iv)
}
