package fileKit

import (
	"os"
	"path/filepath"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

// Remove 删除单个文件或空目录.
/*
@param path 	（1）目标不存在 → 返回错误
				（2）目录非空 → 返回错误
*/
func Remove(path string) (err error) {
	//return gfile.RemoveFile(path)

	if err = os.Remove(path); err != nil {
		err = errKit.Wrapf(err, `os.Remove failed for path "%s"`, path)
	}
	return
}

// RemoveAll 递归删除文件、空目录或非空目录，类似 rm -rf.
/*
@param path 	（1）目标不存在 → 不报错，返回 nil
				（2）路径为空字符串 → 不做任何操作
*/
func RemoveAll(path string) (err error) {
	//return gfile.RemoveAll(path)

	if err = os.RemoveAll(path); err != nil {
		err = errKit.Wrapf(err, `os.RemoveAll failed for path "%s"`, path)
	}
	return
}

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
