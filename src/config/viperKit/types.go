package viperKit

import "github.com/richelieu042/chimera/v3/src/file/fileKit"

type (
	// Data 配置文件的数据.
	Data struct {
		// Content 配置文件的内容
		Content []byte `json:"data" yaml:"data"`

		// Type 配置文件的类型，"yaml", "yml", "json", "toml", "hcl", "tfvars", "dotenv", "env", "properties", "props", "prop", "ini"
		Type string `json:"type" yaml:"type"`
	}
)

func NewDataFromFile(path string) (*Data, error) {
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return nil, err
	}

	content, err := fileKit.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Data{
		Content: content,
		Type:    GetContentType(path),
	}, nil
}
