package fileKit

import (
	"os"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

var (
	// Open （以指定 flag 和 perm）打开文件/目录.
	/*
		@param flag 详见"_info.md".
		@param perm	(1) 可以参考 "fileKit/consts.go"
					(2) e.g.0666 || os.ModePerm ...
	*/
	Open func(path string, flag int, perm os.FileMode) (*os.File, error) = gfile.OpenFile

	// OpenReadOnly （只读模式）打开文件/目录.
	/*
		!!!: 只读权限 打开.
	*/
	OpenReadOnly func(path string) (*os.File, error) = gfile.Open
)

// Create 创建文件（读写权限、文件不存在就创建、打开并清空文件）.
/*
TODO: perm 权限可自定义，看后续 gfile 后不会完善.

@param filePath 会尝试为其创建父目录

PS:
(1) 读写权限（0644）；
(2) path不存在，会创建；
(3) path存在 && 是文件，会 截断 该文件内容；
(4) path存在 && 是目录，会返回error.
*/
func Create(filePath string) (*os.File, error) {
	//return gfile.Create(filePath)

	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := MkParentDirs(filePath); err != nil {
		return nil, err
	}

	return Open(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
}

// CreateInAppendMode 创建文件（读写权限、文件不存在就创建、追加模式）.
/*
TODO: perm 权限可自定义，看后续 gfile 后不会完善.
*/
func CreateInAppendMode(filePath string) (*os.File, error) {
	//return gfile.Create(filePath)

	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := MkParentDirs(filePath); err != nil {
		return nil, err
	}

	return Open(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
}

// NewTemporaryFile 在指定目录下，生成临时文件.
/*
@param dirPath 如果为""，临时文件将生成在 系统临时目录 内；如果为"."，临时文件将生成在 当前目录 内.

e.g.
pattern: "tempfile_test" 		=> 临时文件的文件名: "tempfile_test2594316144"
pattern: "tempfile_test*" 		=> 临时文件的文件名: "tempfile_test827818253"
pattern: "tempfile_test*.xyz" 	=> 临时文件的文件名: "tempfile_test3617672388.xyz"
*/
func NewTemporaryFile(dirPath, pattern string) (*os.File, error) {
	if err := AssertNotExistOrIsDir(dirPath); err != nil {
		return nil, err
	}
	if err := MkDirs(dirPath); err != nil {
		return nil, err
	}

	if err := strKit.AssertNotBlank(pattern, "pattern"); err != nil {
		return nil, err
	}

	return os.CreateTemp(dirPath, pattern)
}
