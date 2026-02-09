package fileKit

import (
	"bytes"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// ReadLuaFileToString 按行读取 .lua文件 的内容.
/*
@param path .lua文件的路径
*/
func ReadLuaFileToString(filePath string) (string, error) {
	buf := bytes.NewBuffer(nil)

	err := ReadLines(filePath, func(line string) error {
		line = strKit.TrimSpace(line)

		// 忽略"空行"和"注释行"
		if strKit.IsEmpty(line) || strKit.StartWith(line, "--") {
			return nil
		}

		buf.WriteString(line)
		// 再加个空格
		buf.WriteString(" ")
		return nil
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
