package fileKit

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

// Predicate 判断文件（or空目录）是否应该被删除的回调函数类型
type Predicate func(path string, info os.FileInfo) bool

// Clean 递归清理路径下满足所有 predicate 条件的文件，并删除空目录.
/*
!!!: 目录的修改时间 ModTime 比较特殊，会因为其内部的操作改变而改变的，e.g. 新增文件/目录、删除文件/目录、重命名文件 / 目录（同一目录下）、移动文件（跨目录）...

@param path 		文件或目录的路径（如果不存在，将返回nil）
@param predicates	（1）所有 predicate 返回 true 时才删除文件或空目录（AND 逻辑）
					（2）如果不传，将整个删除 path 对应的文件或目录
*/
func Clean(path string, predicates ...Predicate) error {
	if len(predicates) == 0 {
		return RemoveAll(path)
	}

	// Lstat 不会跟随根符号链接，避免递归清理链接指向的目录。
	info, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return errKit.Wrapf(err, "fail to lstat path(%s)", path)
	}

	if info.IsDir() {
		dir, err := openVerifiedDir(path)
		if err != nil {
			return errKit.Wrapf(err, "fail to open dir(%s)", path)
		}
		defer dir.Close()

		deleteRoot, err := cleanDirectory(dir.root, path, predicates)
		if err != nil {
			return err
		}
		if deleteRoot {
			if err := dir.Remove(); err != nil && !errors.Is(err, os.ErrNotExist) {
				return errKit.Wrapf(err, "fail to remove dir(%s)", path)
			}
		}
		return nil
	}
	return cleanFile(path, info, predicates)
}

func cleanDirectory(root *os.Root, dirPath string, predicates []Predicate) (bool, error) {
	dir, err := root.Open(".")
	if err != nil {
		return false, errKit.Wrapf(err, "fail to open dir(%s)", dirPath)
	}
	entries, readErr := dir.ReadDir(-1)
	closeErr := dir.Close()
	if readErr != nil {
		return false, errKit.Wrapf(readErr, "fail to read dir(%s)", dirPath)
	}
	if closeErr != nil {
		return false, errKit.Wrapf(closeErr, "fail to close dir(%s)", dirPath)
	}

	for _, entry := range entries {
		name := entry.Name()
		entryPath := filepath.Join(dirPath, name)
		info, err := root.Lstat(name)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return false, errKit.Wrapf(err, "fail to lstat entry(%s)", entryPath)
		}

		if info.IsDir() {
			child, err := root.OpenRoot(name)
			if err != nil {
				return false, errKit.Wrapf(err, "fail to open dir(%s)", entryPath)
			}
			openedInfo, statErr := child.Stat(".")
			if statErr != nil {
				child.Close()
				return false, errKit.Wrapf(statErr, "fail to stat dir(%s)", entryPath)
			}
			if !os.SameFile(info, openedInfo) {
				child.Close()
				return false, errKit.Newf("dir(%s) changed while being opened", entryPath)
			}

			deleteChild, cleanErr := cleanDirectory(child, entryPath, predicates)
			closeErr := child.Close()
			if cleanErr != nil {
				return false, cleanErr
			}
			if closeErr != nil {
				return false, errKit.Wrapf(closeErr, "fail to close dir(%s)", entryPath)
			}
			if deleteChild {
				if err := root.Remove(name); err != nil && !errors.Is(err, os.ErrNotExist) {
					return false, errKit.Wrapf(err, "fail to remove dir(%s)", entryPath)
				}
			}
		} else {
			if err := cleanRootFile(root, name, entryPath, info, predicates); err != nil {
				return false, err
			}
		}
	}

	empty, err := isRootEmpty(root)
	if err != nil {
		return false, errKit.Wrapf(err, "fail to judge if dir(%s) is empty", dirPath)
	}
	if !empty {
		return false, nil
	}

	info, err := root.Stat(".")
	if err != nil {
		return false, errKit.Wrapf(err, "fail to stat dir(%s)", dirPath)
	}
	return canDelete(dirPath, info, predicates), nil
}

func cleanRootFile(root *os.Root, name, path string, info os.FileInfo, predicates []Predicate) error {
	if !canDelete(path, info, predicates) {
		return nil
	}
	if err := root.Remove(name); err != nil && !errors.Is(err, os.ErrNotExist) {
		return errKit.Wrapf(err, "fail to remove file(%s)", path)
	}
	return nil
}

func isRootEmpty(root *os.Root) (bool, error) {
	dir, err := root.Open(".")
	if err != nil {
		return false, err
	}
	entries, readErr := dir.ReadDir(1)
	closeErr := dir.Close()
	if readErr != nil && !errors.Is(readErr, io.EOF) {
		return false, readErr
	}
	if closeErr != nil {
		return false, closeErr
	}
	return len(entries) == 0, nil
}

func cleanFile(filePath string, info os.FileInfo, predicates []Predicate) error {
	if !canDelete(filePath, info, predicates) {
		return nil // predicates 不允许删除
	}

	// 安全删除一个可能不存在的文件
	if err := Remove(filePath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return errKit.Wrapf(err, "fail to remove file(%s)", filePath)
	}

	/* 成功删除：文件 */
	return nil
}

/*
@param path 文件(或空目录)的路径
*/
func canDelete(path string, info os.FileInfo, predicates []Predicate) bool {
	for _, predicate := range predicates {
		if !predicate(path, info) {
			return false
		}
	}
	return true
}
