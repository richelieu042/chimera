// Package aesKit
/*
AES/ECB/PKCS5Padding
参考：https://www.cnblogs.com/lavin/p/5373188.html
缺点：代码中有 panic
秘钥：长度必须是 16、24、32 中的一个！
*/
package aesKit1

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
)

// EncryptToString 加密流程: 明文 => 密文 => base64字符串
/*
@param plainText 可以为""
*/
func EncryptToString(plainText, key []byte) (string, error) {
	cipherText, err := Encrypt(plainText, key)
	if err != nil {
		return "", err
	}
	return base64Kit.EncodeToString(cipherText), nil
}

func Encrypt(plainText, key []byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		err = errorKit.Wrapf(err, "key(%s) is invalid", string(key))
		return
	}

	defer func() {
		if obj := recover(); obj != nil {
			cipherText = nil
			err = errorKit.Newf(strKit.ToString(obj))
		}
	}()

	ecb := NewECBEncrypter(block)
	content := PKCS5Padding(plainText, block.BlockSize())
	cipherText = make([]byte, len(content))
	ecb.CryptBlocks(cipherText, content)
	return
}

func DecryptToString(base64Text, key []byte) (string, error) {
	cipherText, err := base64Kit.Decode(base64Text)
	if err != nil {
		return "", err
	}

	plainText, err := Decrypt(cipherText, key)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func Decrypt(cipherText, key []byte) (plainText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		err = errorKit.Wrapf(err, "key(%s) is invalid", string(key))
		return
	}

	defer func() {
		if obj := recover(); obj != nil {
			plainText = nil
			err = errorKit.Newf(strKit.ToString(obj))
		}
	}()

	blockMode := NewECBDecrypter(block)
	plainText = make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = PKCS5Unpadding(plainText)
	return
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int {
	return x.blockSize
}

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
