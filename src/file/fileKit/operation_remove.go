package fileKit

import (
	"errors"
	"os"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

// Remove 删除单个文件或空目录.
/*
使用建议：明确只删单文件时用 Remove（更安全，避免误删）；需要清理目录树时才用 RemoveAll。

@param path 	（1）目标不存在 → 返回错误，可以通过 os.IsNotExist(err) 判断出
				（2）目录非空 → 返回错误

e.g. 安全删除一个可能不存在的文件
if err := os.Remove("tmp.txt"); err != nil && !os.IsNotExist(err) {
	log.Fatal(err)
}
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
使用建议：明确只删单文件时用 Remove（更安全，避免误删）；需要清理目录树时才用 RemoveAll。

@param path 	（1）目标不存在 → 不报错，返回 nil
				（2）路径为空字符串 → 不做任何操作

e.g. 清理整个临时目录（无论是否存在）
if err := os.RemoveAll("./tmp"); err != nil {
    log.Fatal(err)
}
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
@param dirPath 	（1）可以不存在（此时将返回nil）
			（2）不能是符号链接，避免清空链接指向的目录
*/
func EmptyDir(dirPath string) error {
	dir, err := openVerifiedDir(dirPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return errKit.Wrapf(err, "fail to open dir(%s)", dirPath)
	}
	defer dir.Close()

	rootFile, err := dir.root.Open(".")
	if err != nil {
		return errKit.Wrapf(err, "fail to open dir(%s)", dirPath)
	}
	dirEntries, readErr := rootFile.ReadDir(-1)
	closeErr := rootFile.Close()
	if readErr != nil {
		return errKit.Wrapf(readErr, "fail to read dir(%s)", dirPath)
	}
	if closeErr != nil {
		return errKit.Wrapf(closeErr, "fail to close dir(%s)", dirPath)
	}

	for _, dirEntry := range dirEntries {
		if err := dir.root.RemoveAll(dirEntry.Name()); err != nil {
			return errKit.Wrapf(err, "fail to remove entry(%s) from dir(%s)", dirEntry.Name(), dirPath)
		}
	}
	return nil
}
