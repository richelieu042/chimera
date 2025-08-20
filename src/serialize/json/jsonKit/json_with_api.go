package jsonKit

import (
	"os"

	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

func MarshalWithAPI(api API, v interface{}) ([]byte, error) {
	if api == nil {
		api = defaultApi
	}

	return api.Marshal(v)
}

// MarshalIndentWithAPI
/*
@param indent 为了兼容性，用"    "（4个空格）替代"\t"
*/
func MarshalIndentWithAPI(api API, v interface{}, prefix, indent string) ([]byte, error) {
	if api == nil {
		api = defaultApi
	}

	return api.MarshalIndent(v, prefix, indent)
}

func MarshalToStringWithAPI(api API, v interface{}) (string, error) {
	if api == nil {
		api = defaultApi
	}

	return api.MarshalToString(v)
}

// MarshalIndentToStringWithAPI
/*
@param indent 为了兼容性，用"    "（4个空格）替代"\t"
*/
func MarshalIndentToStringWithAPI(api API, v interface{}, prefix, indent string) (string, error) {
	data, err := MarshalIndentWithAPI(api, v, prefix, indent)
	return string(data), err
}

// MarshalToFileWithAPI
/*
PS:
(1) 对 传参filePath 的验证和断言在 fileKit.WriteToFile 里面;
(2) 写入文件的json带有 indent.

@param in		建议为结构体实例指针 || map实例 || slice实例
@param filePath (1) .json 格式的文件
				(2) 不存在的话，会创建一个新的文件
				(3) 存在且是个文件的话，会 "覆盖" 掉旧的（并不会加到该文件的最后面）
*/
func MarshalToFileWithAPI(api API, in interface{}, filePath string, perm os.FileMode) error {
	data, err := MarshalIndentWithAPI(api, in, "", "    ")
	if err != nil {
		return err
	}
	return fileKit.WriteToFile(filePath, data)
}

func UnmarshalWithAPI(api API, data []byte, v interface{}) error {
	if api == nil {
		api = defaultApi
	}

	return api.Unmarshal(data, v)
}

func UnmarshalFromStringWithAPI(api API, str string, v interface{}) error {
	if api == nil {
		api = defaultApi
	}

	return api.UnmarshalFromString(str, v)
}
