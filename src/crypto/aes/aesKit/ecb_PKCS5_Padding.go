package aesKit

//import (
//	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
//	"github.com/zeromicro/go-zero/core/codec"
//)
//
//// EncryptEcbPKCS5PaddingToString 明文([]byte) => 密文(base64 string)
///*
//@param key 	密钥（must be 16/24/32 bit length）
//*/
//func EncryptEcbPKCS5PaddingToString(key []byte, plainData []byte) (string, error) {
//	cipherData, err := EncryptEcbPKCS5Padding(key, plainData)
//	if err != nil {
//		return "", err
//	}
//	return base64Kit.EncodeToString(cipherData), nil
//}
//
//// DecryptEcbPKCS5PaddingFromString 密文(base64 string) => 明文([]byte)
///*
//@param key 	密钥（must be 16/24/32 bit length）
//*/
//func DecryptEcbPKCS5PaddingFromString(key []byte, cipherText string) ([]byte, error) {
//	cipherData, err := base64Kit.DecodeString(cipherText)
//	if err != nil {
//		return nil, err
//	}
//	return DecryptEcbPKCS5Padding(key, cipherData)
//}
//
//// EncryptEcbPKCS5Padding 明文([]byte) => 密文([]byte)
//var EncryptEcbPKCS5Padding func(key, plainText []byte) ([]byte, error) = codec.EcbEncrypt
//
//// DecryptEcbPKCS5Padding 密文([]byte) => 明文([]byte)
//var DecryptEcbPKCS5Padding func(key, cipherData []byte) ([]byte, error) = codec.EcbDecrypt
