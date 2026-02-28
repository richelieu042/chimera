package rsaKit

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// GenerateKeyFiles 生成: 公钥 && 私钥
/*
@param bits		512 ||1024 || 2048 || 3072 || 4096
@param format	PKCS1 || PKCS8
@param password 私钥的密码（没有则传""）
@param priPath	生成私钥文件的位置
@param pubPath	生成公钥文件的位置
*/
func GenerateKeyFiles(bits int, format KeyFormat, password, priPath, pubPath string, perm os.FileMode) error {
	pri, pub, err := GenerateKeys(bits, format, password)
	if err != nil {
		return err
	}

	if err := fileKit.WriteToFile(priPath, pri); err != nil {
		return err
	}
	if err := fileKit.WriteToFile(pubPath, pub); err != nil {
		return err
	}
	return nil
}

func GenerateKeys(bits int, format KeyFormat, password string) (pri []byte, pub []byte, err error) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	var derStream []byte
	switch format {
	case PKCS1:
		derStream = x509.MarshalPKCS1PrivateKey(privateKey)
	case PKCS8:
		derStream, err = x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errorKit.Newf("invalid format(%v)", format)
	}
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	pri = pem.EncodeToMemory(block)
	pri, err = EncryptPrivatePEM(pri, format, password)
	if err != nil {
		return nil, nil, err
	}

	// 生成公钥
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pub = pem.EncodeToMemory(block)

	return pri, pub, nil
}
