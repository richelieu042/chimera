package pathKit

import (
	"path/filepath"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

var (
	// RealPath 获取给定路径的绝对路径地址.
	/*
		PS: 对应的文件或目录不存在的话，返回"".

		e.g.
		fmt.Println(pathKit.RealPath(""))   // "/Users/richelieu/GolandProjects/chimera"
		fmt.Println(pathKit.RealPath("."))  // "/Users/richelieu/GolandProjects/chimera"
		fmt.Println(pathKit.RealPath("./")) // "/Users/richelieu/GolandProjects/chimera"
	*/
	RealPath func(path string) string = gfile.RealPath

	// SelfDir returns absolute directory path of current running process(binary).
	/*
		e.g.
		fmt.Println(pathKit.SelfDir()) // "/Users/richelieu/Library/Caches/JetBrains/GoLand2023.3/tmp/GoLand"
	*/
	SelfDir func() string = gfile.SelfDir

	// MainPkgPath returns absolute file path of package main, which contains the entrance function main.
	MainPkgPath func() string = gfile.MainPkgPath

	// ParentDir 返回文件（或目录）路径的父路径（父目录）.
	/*
	   PS: 类似于Java中的 getParentFile().
	*/
	ParentDir func(path string) string = gfile.Dir

	// IsAbs 判断路径是不是绝对路径
	/*
	   PS: 传参path 对应的文件（或目录）可以不存在.

	   e.g.
	   ("./a/b/c") 	=> false
	   ("C:/a/b/c") 	=> true
	   ("/root") 		=> true
	   ("root") 		=> false
	*/
	IsAbs func(path string) bool = filepath.IsAbs

	// Abs 绝对路径.
	/*
		@return 传参path相对于当前路径(os.Getwd())
	*/
	Abs func(path string) (string, error) = filepath.Abs

	// VolumeName 返回分区名
	/*
	   PS: 非Windows平台，将返回"".

	   e.g. Windows平台
	   	("C:/a/b/c") => "C:"
	*/
	VolumeName func(path string) string = filepath.VolumeName

	// GetRelativePath 获取（targetPath 相对于 basePath 的）相对路径.
	/*
	   e.g. Windows
	   ("C:/a/b", "C:/a/b/c/d/../e") => "c\e", <nil>

	   e.g.1 Mac
	   ("/usr/local", "/usr/local/go/bin")			=> "go/bin", nil
	   ("//usr////local", "/usr/local/go/bin")		=> "go/bin", nil
	   ("//usr////local", "/usr/local/go/bin/../")	=> "go", nil
	*/
	GetRelativePath func(basepath, targpath string) (string, error) = filepath.Rel

	// Clean 返回等价的最短路径（清理路径）
	/*
	   PS: 可用于去掉路径中的 "./" 、 "../" ...

	   e.g.
	   ("./1.txt") 		=> "1.txt"
	   ("/root/.././c") 	=> "/c"
	*/
	Clean func(path string) string = filepath.Clean

	// Match （正则相关）匹配文件名，完全匹配则返回true
	/*
	   e.g.
	   ("*", "a") => (true, <nil>)
	   ("*", "C:/a/b/c") => (true, <nil>)
	   ("\\b", "b") => (false, <nil>)
	*/
	Match func(pattern, name string) (matched bool, err error) = filepath.Match

	// Glob （正则相关）返回所有匹配的文件.
	/*
	   e.g.
	   ("d:/test/*.txt") => ([d:\test\a.txt d:\test\b.txt], <nil>)
	*/
	Glob func(pattern string) (matches []string, err error) = filepath.Glob

	// EvalSymlinks 返回链接文件的实际路径.
	EvalSymlinks func(path string) (string, error) = filepath.EvalSymlinks
)

// CheckSkip 检查是否发生"路径穿越"（路径穿透）.
/*
@param path	 	父路径
@param path1	子路径
@return true: 发生"路径穿越"

e.g.
("/a//b/", "/a//b//../c.docx")	=> true
("/a//b/", "//a//b///c.docx")	=> false
*/
func CheckSkip(parent, path string) bool {
	parent = Join(parent)
	path = Join(path)

	return !strKit.StartWith(path, parent)
}
