package fileKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

// AssertExist
/*
@param path 文件（或目录）的路径
*/
func AssertExist(path string) error {
	if strKit.IsBlank(path) {
		if path == "" {
			return errorKit.NewfWithSkip(1, "[%s] path is empty", funcKit.GetFuncName(1))
		}
		return errorKit.NewfWithSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertNotExistOrIsFile
/*
通过的情况: 	不存在 || 存在但是个文件
不通过的情况:	存在但是个目录

@param mkdirArgs 	(1) true: path不存在的话，为其创建父目录（可多级）
					(2) 默认: true
*/
func AssertNotExistOrIsFile(path string) error {
	if strKit.IsBlank(path) {
		if path == "" {
			return errorKit.NewfWithSkip(1, "[%s] path is empty", funcKit.GetFuncName(1))
		}
		return errorKit.NewfWithSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if Exists(path) && IsDir(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) exists but it is a directory", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertNotExistOrIsDir
/*
通过的情况: 	不存在 || 存在但是个目录
不通过的情况:	存在但是个文件

@param mkdirArgs	(1) true: path不存在的话，为其创建目录（可多级）
					(2) 默认: true
*/
func AssertNotExistOrIsDir(path string) error {
	if strKit.IsBlank(path) {
		if path == "" {
			return errorKit.NewfWithSkip(1, "[%s] path is empty", funcKit.GetFuncName(1))
		}
		return errorKit.NewfWithSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if Exists(path) && IsFile(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) exists but it is a file", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertExistAndIsFile
/*
@return 如果path存在且是个文件，返回nil
*/
func AssertExistAndIsFile(path string) error {
	if strKit.IsBlank(path) {
		if path == "" {
			return errorKit.NewfWithSkip(1, "[%s] path is empty", funcKit.GetFuncName(1))
		}
		return errorKit.NewfWithSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if IsDir(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) exists but it is a directory", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertExistAndIsDir
/*
@return 如果path存在且是个目录，返回nil
*/
func AssertExistAndIsDir(path string) error {
	if strKit.IsBlank(path) {
		if path == "" {
			return errorKit.NewfWithSkip(1, "[%s] path is empty", funcKit.GetFuncName(1))
		}
		return errorKit.NewfWithSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if IsFile(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) exists but it is a file", funcKit.GetFuncName(1), path)
	}
	return nil
}

func AssertReadableAndWritable(path string) error {
	if strKit.IsBlank(path) {
		if path == "" {
			return errorKit.NewfWithSkip(1, "[%s] path is empty", funcKit.GetFuncName(1))
		}
		return errorKit.NewfWithSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if !IsReadable(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) isn't readable", funcKit.GetFuncName(1), path)
	}
	if !IsWritable(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) isn't writable", funcKit.GetFuncName(1), path)
	}
	return nil
}

func AssertReadable(path string) error {
	if strKit.IsBlank(path) {
		if path == "" {
			return errorKit.NewfWithSkip(1, "[%s] path is empty", funcKit.GetFuncName(1))
		}
		return errorKit.NewfWithSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if !IsReadable(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) isn't readable", funcKit.GetFuncName(1), path)
	}
	return nil
}

func AssertWritable(path string) error {
	if strKit.IsBlank(path) {
		if path == "" {
			return errorKit.NewfWithSkip(1, "[%s] path is empty", funcKit.GetFuncName(1))
		}
		return errorKit.NewfWithSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if !IsWritable(path) {
		return errorKit.NewfWithSkip(1, "[%s] path(%s) isn't writable", funcKit.GetFuncName(1), path)
	}
	return nil
}
