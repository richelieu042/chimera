package fileKit

import (
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/os/gfile"
)

var (
	// RemoveFile 删除文件.
	RemoveFile func(path string) (err error) = gfile.RemoveFile

	// RemoveAll 删除文件（或目录）.
	/*
		PS: 如果是目录且内部有文件或目录，也会一并删除.
	*/
	RemoveAll func(path string) (err error) = gfile.RemoveAll
)

// EmptyDir 清空目录：删掉目录中的文件和子目录（递归），但该目录本身不会被删掉.
/*
@param dirPath 可以不存在（此时将返回nil）
*/
func EmptyDir(dirPath string) error {
	if !Exists(dirPath) {
		return nil
	}
	if err := AssertExistAndIsDir(dirPath); err != nil {
		return err
	}

	// 遍历目录
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, dirEntry := range dirEntries {
		path := filepath.Join(dirPath, dirEntry.Name())
		if err := RemoveAll(path); err != nil {
			return err
		}
	}
	return nil
}
