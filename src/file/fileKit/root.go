package fileKit

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

var (
	// OpenRoot
	/*
		https://cloud.tencent.com/developer/article/2522181
		https://cloud.tencent.com/developer/article/2502915

		os.Root 可以锁定工作目录。 使用户无法打开目录外的文件，例如 ../../../etc/passwd 。
		可以算一种 安全保护，最重要的是 强制约束用户， 限制用户行为， 检查计划外的使用逻辑 。 免得和煞笔瞎掰扯， 浪费时间。
	*/
	OpenRoot func(dir string) (*os.Root, error) = os.OpenRoot

	OpenInRoot func(dir, name string) (*os.File, error) = os.OpenInRoot
)

type verifiedDir struct {
	root   *os.Root
	parent *os.Root
	name   string
	path   string
}

func (dir *verifiedDir) Close() error {
	if dir.parent == nil {
		return dir.root.Close()
	}
	return errors.Join(dir.root.Close(), dir.parent.Close())
}

func (dir *verifiedDir) Remove() error {
	if dir.parent == nil {
		return os.Remove(dir.path)
	}
	return dir.parent.Remove(dir.name)
}

// openVerifiedDir opens path as an anchored directory and rejects symbolic links.
// The identity check closes the race between inspecting path and opening it.
func openVerifiedDir(path string) (*verifiedDir, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, errKit.Wrapf(err, "fail to get absolute path of dir(%s)", path)
	}
	absPath = filepath.Clean(absPath)
	parentPath := filepath.Dir(absPath)
	name := filepath.Base(absPath)

	if parentPath == absPath {
		info, err := os.Lstat(absPath)
		if err != nil {
			return nil, err
		}
		root, err := openAndVerifyDir(absPath, info)
		if err != nil {
			return nil, err
		}
		return &verifiedDir{root: root, path: absPath}, nil
	}

	parent, err := os.OpenRoot(parentPath)
	if err != nil {
		return nil, err
	}

	info, err := parent.Lstat(name)
	if err != nil {
		parent.Close()
		return nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		parent.Close()
		return nil, errKit.Newf("dir(%s) is a symbolic link", path)
	}
	if !info.IsDir() {
		parent.Close()
		return nil, errKit.Newf("path(%s) exists but it is not a directory", path)
	}

	root, err := parent.OpenRoot(name)
	if err != nil {
		parent.Close()
		return nil, err
	}
	openedInfo, err := root.Stat(".")
	if err != nil {
		root.Close()
		parent.Close()
		return nil, err
	}
	if !os.SameFile(info, openedInfo) {
		root.Close()
		parent.Close()
		return nil, errKit.Newf("dir(%s) changed while being opened", path)
	}
	return &verifiedDir{root: root, parent: parent, name: name, path: absPath}, nil
}

func openAndVerifyDir(path string, info os.FileInfo) (*os.Root, error) {
	if info.Mode()&os.ModeSymlink != 0 {
		return nil, errKit.Newf("dir(%s) is a symbolic link", path)
	}
	if !info.IsDir() {
		return nil, errKit.Newf("path(%s) exists but it is not a directory", path)
	}

	root, err := os.OpenRoot(path)
	if err != nil {
		return nil, err
	}
	openedInfo, err := root.Stat(".")
	if err != nil {
		root.Close()
		return nil, err
	}
	if !os.SameFile(info, openedInfo) {
		root.Close()
		return nil, errKit.Newf("dir(%s) changed while being opened", path)
	}
	return root, nil
}
