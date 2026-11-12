package yamlKit

import (
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"go.yaml.in/yaml/v4"
)

// Marshal
/*
PS: 需要搭配 yaml tag 一起使用，不识别 json tag.
	e.g.字段 UserName：
	- 仅写 json:"user_name" → YAML 键名为 username。
	- 加上 yaml:"user_name" → YAML 键名才是 user_name。
*/
var Marshal func(in any) (out []byte, err error) = yaml.Marshal

func MarshalToFile(in any, filePath string) error {
	data, err := Marshal(in)
	if err != nil {
		return err
	}
	return fileKit.WriteToFile(filePath, data)
}
