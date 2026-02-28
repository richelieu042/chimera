package fileKit

import (
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// Move 移动.
/*
PS: 移动成功后，path 将不再存在.
*/
var Move func(src, dst string) error = Rename

// Rename 重命名.
/*
PS:
(1) 参考 gfile.Rename;
(2) 重命名成功后，src 将不再存在.
*/
func Rename(src, dst string) error {
	if err := AssertExist(src); err != nil {
		return err
	}
	if err := MkParentDirs(dst); err != nil {
		return err
	}

	err := os.Rename(src, dst)
	if err != nil {
		err = errorKit.Wrapf(err, `fail to rename from "%s" to "%s"`, src, dst)
	}
	return err
}

// RenameInSameDir 同目录下重命名.
/*
PS: 重命名成功后，path 将不再存在.
*/
func RenameInSameDir(path string, name string) error {
	if err := AssertExist(path); err != nil {
		return err
	}
	if err := strKit.AssertNotBlank(name, "name"); err != nil {
		return err
	}

	parentDir := gfile.Dir(path)
	return Rename(path, filepath.Join(parentDir, name))
}
