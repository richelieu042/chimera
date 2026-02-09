package jsonRespKit

import (
	"github.com/richelieu042/chimera/v3/src/config/viperKit"
	"github.com/richelieu042/chimera/v3/src/core/mapKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

type FileData struct {
	// Type 文件的类型，e.g. "properties"、"yaml"、"ini"...
	Type string

	// Data 文件的内容
	Data []byte
}

/*
key:	code
value:	message
*/
var msgMap = make(map[string]string)

// readFile 读取message文件，加到 msgMap 中.
/*
@param filePath 建议是 .properties 的文件
*/
func readFile(filePath string) error {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return err
	}

	m := make(map[string]string)
	if _, err := viperKit.UnmarshalFromFile(filePath, nil, &m); err != nil {
		return err
	}
	msgMap = mapKit.Merge(msgMap, m)
	return nil
}

func readFileData(fd *FileData) error {
	if fd == nil {
		return nil
	}

	m := make(map[string]string)
	if _, err := viperKit.Unmarshal(fd.Data, fd.Type, nil, &m); err != nil {
		return err
	}
	msgMap = mapKit.Merge(msgMap, m)
	return nil
}
