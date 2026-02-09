package rsaKit

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// ParsePublicKeyFromPem 解析公钥.
/*
支持: PKCS1、PKCS8.
*/
func ParsePublicKeyFromPem(pemData []byte) (*rsa.PublicKey, error) {
	return jwt.ParseRSAPublicKeyFromPEM(pemData)

	//block, _ := pem.Decode(pemData)
	//if block == nil {
	//	return nil, errorKit.Sign("fail to decode pem because block is nil")
	//}
	//
	//keyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	//if err != nil {
	//	return nil, err
	//}
	//return keyInterface.(*rsa.PublicKey), nil
}

// ParsePrivateKeyFromPem 解析私钥.
/*
支持: PKCS1、PKCS8.

@param password 私钥的密码（"": 无密码）
*/
func ParsePrivateKeyFromPem(data []byte, password string) (*rsa.PrivateKey, error) {
	if strKit.IsNotEmpty(password) {
		return jwt.ParseRSAPrivateKeyFromPEMWithPassword(data, password)
	}
	return jwt.ParseRSAPrivateKeyFromPEM(data)

	//block, _ := pem.Decode(data)
	//if block == nil {
	//	return nil, errorKit.Sign("fail to decode pem because block is nil")
	//}
	//
	//switch opts.format {
	//case PKCS1:
	//	return x509.ParsePKCS1PrivateKey(block.Bytes)
	//case PKCS8:
	//	keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return keyInterface.(*rsa.PrivateKey), nil
	//default:
	//	return nil, errorKit.Sign("invalid key format(%d)", opts.format)
	//}
}
