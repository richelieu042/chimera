package fileKit

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

var (
	Exists func(path string) bool = gfile.Exists

	// IsFile 如果传参path不存在，将返回false.
	IsFile func(path string) bool = gfile.IsFile

	// IsDir 如果传参path不存在，将返回false.
	IsDir func(path string) bool = gfile.IsDir

	// Stat 获取文件（或目录）信息
	/*
		@param path 如果为""或不存在，将返回error(e.g."" => stat : no such file or directory)
	*/
	Stat func(path string) (os.FileInfo, error) = gfile.Stat

	// IsEmpty 传参path对应的文件（或目录）是否为空？
	/*
		If `path` is a folder, it checks if there's any file under it.
		If `path` is a file, it checks if the file size is zero.
		Note that it returns true if `path` does not exist.
	*/
	IsEmpty func(path string) bool = gfile.IsEmpty

	// GetFileName 获取 文件名.
	/*
		e.g.
		/var/www/file.js -> file.js
		file.js          -> file.js
	*/
	GetFileName func(path string) string = filepath.Base

	// GetName 获取 文件名的前缀.
	/*
		e.g.
		/var/www/file.js -> file
		file.js          -> file
	*/
	GetName func(path string) string = gfile.Name
)

// GetExt 获取 文件名的后缀（带"."）
/*
	@param path 		可以不存在（exist）
	@param lowerArgs	是否将返回值转换为小写字母? 默认true
	@return 可能为""

	e.g.
		println(fileKit.GetExt("main.go"))  // ".go"
		println(fileKit.GetExt("api.json")) // ".json"
		println(fileKit.GetExt(""))         // ""
		println(fileKit.GetExt("    "))     // ""
		println(fileKit.GetExt("empty"))    // ""
	e.g.1
		("./iShot_2024-09-04_13.51.58.PNG") => ".png"
*/
func GetExt(path string, lowerArgs ...bool) (ext string) {
	path = strKit.TrimSpace(path)
	ext = gfile.Ext(path)
	if lowerArgs == nil || lowerArgs[0] {
		// 手动转换为小写字母
		ext = strings.ToLower(ext)
	}
	return
}

// GetExtName 获取后缀（不带"."）.
/*
	@param path 		可以不存在（exist）
	@param lowerArgs	是否将返回值转换为小写字母? 默认true
	@return 可能为""

	e.g.
		println(fileKit.GetExtName("main.go"))  // "go"
		println(fileKit.GetExtName("api.json")) // "json"
		println(fileKit.GetExtName(""))         // ""
		println(fileKit.GetExtName("    "))     // ""
		println(fileKit.GetExtName("empty"))    // ""
	e.g.1
		("./iShot_2024-09-04_13.51.58.PNG") => "png"
*/
func GetExtName(path string, lowerArgs ...bool) (extName string) {
	extName = gfile.ExtName(path)

	if lowerArgs == nil || lowerArgs[0] {
		// 手动转换为小写字母
		extName = strings.ToLower(extName)
	}

	return
}
