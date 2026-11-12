package yamlKit

import (
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"go.yaml.in/yaml/v4"
)

// Unmarshal
/*
PS: 需要搭配 yaml tag 一起使用，不识别 json tag.
*/
var Unmarshal func(in []byte, out any) (err error) = yaml.Unmarshal

func UnmarshalFromFile(path string, out any) error {
	data, err := fileKit.ReadFile(path)
	if err != nil {
		return err
	}

	return Unmarshal(data, out)
}
