package yamlKit

import (
	"os"

	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"gopkg.in/yaml.v3"
)

// Marshal
/*
PS: 需要搭配 yaml tag 一起使用，不识别 json tag.
*/
var Marshal func(in interface{}) (out []byte, err error) = yaml.Marshal

// MarshalToString
/*
PS: 需要搭配 yaml tag 一起使用，不识别 json tag.

@param in 建议为结构体实例指针 || map实例 || slice实例
*/
func MarshalToString(in interface{}) (string, error) {
	data, err := Marshal(in)
	return string(data), err
}

// MarshalToFile
/*
PS:
(1) 需要搭配 yaml tag 一起使用，不识别 json tag.
(2) 对 传参filePath 的验证和断言在 fileKit.WriteToFile 里面.

@param in		建议为结构体实例指针 || map实例 || slice实例
@param filePath (1) .yaml 格式的文件
				(2) 不存在的话，会创建一个新的文件
				(3) 存在且是个文件的话，会 "覆盖" 掉旧的（并不会加到该文件的最后面）
*/
func MarshalToFile(in interface{}, filePath string, perm os.FileMode) error {
	data, err := Marshal(in)
	if err != nil {
		return err
	}
	return fileKit.WriteToFile(filePath, data)
}
