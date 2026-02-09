package sha1Kit

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

var Encrypt func(v interface{}) string = gsha1.Encrypt

// EncryptFile
/*
e.g.
	("/Users/richelieu/Documents/ino/notes/Linux（Unix、Mac）/命令（Mac、Linux）.wps") => "167a518a8ef373c3614f756bd718b27ae8c07853", nil
*/
func EncryptFile(filePath string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return "", err
	}

	return gsha1.EncryptFile(filePath)
}

// EncryptData
/*
参考: gsha1.Encrypt.
*/
func EncryptData(data []byte) string {
	r := sha1.Sum(data)
	return hex.EncodeToString(r[:])
}
