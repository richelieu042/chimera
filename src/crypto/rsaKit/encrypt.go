package rsaKit

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"

	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
)

// Encrypt 公钥加密.
/*
支持: PKCS1、PKCS8.
*/
func Encrypt(data, pemData []byte) ([]byte, error) {
	if err := sliceKit.AssertNotEmpty(pemData, "pemData"); err != nil {
		return nil, err
	}

	// 公钥
	publicKey, err := ParsePublicKeyFromPem(pemData)
	if err != nil {
		return nil, err
	}

	// 分块加密
	buffer := bytes.NewBufferString("")
	partLen := publicKey.N.BitLen()/8 - 11
	chunks := sliceKit.Split(data, partLen)
	for _, chunk := range chunks {
		tmp, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(tmp)
	}
	data = buffer.Bytes()

	// base64编码
	return base64Kit.Encode(data), nil
}
