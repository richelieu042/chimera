package md5Kit

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// Encrypt
/*
Deprecated: 不推荐使用此法，应该明确 传参data 的类型.
*/
var Encrypt func(data interface{}) (string, error) = gmd5.Encrypt

var EncryptBytes func(data []byte) (string, error) = gmd5.EncryptBytes

var EncryptString func(data string) (string, error) = gmd5.EncryptString

// EncryptFile
/*
e.g.
("/Users/richelieu/Documents/ino/notes/Linux（Unix、Mac）/命令（Mac、Linux）.wps") => "390c5cb9447130fa6f42d488630bb459", nil
*/
func EncryptFile(filePath string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return "", err
	}
	return gmd5.EncryptFile(filePath)
}

//// GetFromFile 获取文件的md5
//func GetFromFile(path string) (string, error) {
//	data, err := fileKit.ReadFile(path)
//	if err != nil {
//		return "", err
//	}
//
//	return Get(data), nil
//}
//
//// GetFromFile1
//// Deprecated
///*
//path为 空字符串 或者 对应文件不存在 时，会报错：open : The system cannot find the file specified.
//*/
//func GetFromFile1(path string) (string, error) {
//	file, err := os.OpenReadOnly(path)
//	if err != nil {
//		return "", err
//	}
//
//	m := md5.Sign()
//	_, err = io.Copy(m, file)
//	if err != nil {
//		return "", err
//	}
//	return hex.EncodeToString(m.Sum(nil)), nil
//}
//
//func Get(s []byte) string {
//	m := md5.Sign()
//	m.Write(s)
//	return hex.EncodeToString(m.Sum(nil))
//}
