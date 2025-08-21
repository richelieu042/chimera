package jsonKit

import (
	"os"

	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

func Marshal(v interface{}) ([]byte, error) {
	return defaultApi.Marshal(v)
}

// MarshalIndent
/*
@param indent 为了兼容性，用"    "（4个空格）替代"\t"
*/
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return defaultApi.MarshalIndent(v, prefix, indent)
}

// MarshalToString
/*
e.g.
	fmt.Println(jsonKit.MarshalToString(nil)) // null <nil>
*/
func MarshalToString(v interface{}) (string, error) {
	return defaultApi.MarshalToString(v)
}

// MarshalToFile
/*
PS:
(1) 对 传参filePath 的验证和断言在 fileKit.WriteToFile 里面;
(2) 写入文件的json带有 indent.

@param in		建议为结构体实例指针 || map实例 || slice实例
@param filePath (1) .json 格式的文件
				(2) 不存在的话，会创建一个新的文件
				(3) 存在且是个文件的话，会 "覆盖" 掉旧的（并不会加到该文件的最后面）
*/
func MarshalToFile(in interface{}, filePath string, perm os.FileMode) error {
	data, err := MarshalIndent(in, "", "    ")
	if err != nil {
		return err
	}
	return fileKit.WriteToFile(filePath, data, perm)
}

// MarshalIndentToString
/*
@param indent 为了兼容性，用"    "（4个空格）替代"\t"
*/
func MarshalIndentToString(v interface{}, prefix, indent string) (string, error) {
	data, err := MarshalIndent(v, prefix, indent)
	return string(data), err
}

// Unmarshal 反序列化
/*
@param v	(1) 必须是指针（结构体实例指针 || map实例指针 || slice实例指针）;
			(2) 如果为 slice 或 map 类型，值可以为nil.
*/
func Unmarshal(data []byte, v interface{}) error {
	return defaultApi.Unmarshal(data, v)
}

func UnmarshalFromString(str string, v interface{}) error {
	return defaultApi.UnmarshalFromString(str, v)
}
