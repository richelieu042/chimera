package rsaKit

import (
	"fmt"
	"testing"
)

// 测试: 加密 && 解密
func TestRSA(t *testing.T) {
	bits := 2048
	format := PKCS8
	password := "qwdqwdqwd"
	data := []byte("`1245`152678`16983强无敌群无多7~!@#$%^&*()_+=...￥$")

	// 生成公钥私钥
	pri, pub, err := GenerateKeys(bits, format, password)
	if err != nil {
		panic(err)
	}

	// 加密
	data, err = Encrypt(data, pub)
	if err != nil {
		panic(err)
	}
	fmt.Printf("密文:\n%s\n", string(data))

	// 解密
	data, err = Decrypt(data, pri, password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("明文:\n%s\n", string(data))
}
