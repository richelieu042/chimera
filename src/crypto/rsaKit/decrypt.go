package rsaKit

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"

	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
)

// Decrypt 私钥解密
/*
支持: PKCS1、PKCS8.
*/
func Decrypt(data, pemData []byte, password string) ([]byte, error) {
	if err := sliceKit.AssertNotEmpty(pemData, "pemData"); err != nil {
		return nil, err
	}

	// 私钥
	privateKey, err := ParsePrivateKeyFromPem(pemData, password)
	if err != nil {
		return nil, err
	}

	// base64解码
	data, err = base64Kit.Decode(data)
	if err != nil {
		return nil, err
	}

	// 分块解密
	buffer := bytes.NewBufferString("")
	partLen := privateKey.N.BitLen() / 8
	chunks := sliceKit.Split(data, partLen)
	for _, chunk := range chunks {
		tmp, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(tmp)
	}
	return buffer.Bytes(), nil
}
