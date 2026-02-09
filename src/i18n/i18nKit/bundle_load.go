package i18nKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// LoadMessageFile
/*
@param path 文件的路径（绝对 || 相对; 必须存在且是个文件）
*/
func LoadMessageFile(bundle *i18n.Bundle, path string) (*i18n.MessageFile, error) {
	if err := interfaceKit.AssertNotNil(bundle, "bundle"); err != nil {
		return nil, err
	}
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return nil, err
	}

	return bundle.LoadMessageFile(path)
}

// LoadMessageFileBytes
/*
@param path 文件名 或 文件的路径（绝对 || 相对; 可以不存在）
*/
func LoadMessageFileBytes(bundle *i18n.Bundle, content []byte, path string) (*i18n.MessageFile, error) {
	if err := interfaceKit.AssertNotNil(bundle, "bundle"); err != nil {
		return nil, err
	}
	if err := strKit.AssertNotEmpty(path, "path"); err != nil {
		return nil, err
	}

	return bundle.ParseMessageFileBytes(content, path)
}
