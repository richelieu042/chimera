package pathKit

import (
	"github.com/mitchellh/go-homedir"
	"github.com/richelieu042/chimera/v3/src/core/userKit"
)

// GetHomeDir 获取当前用户的home dir（当前用户主目录）
func GetHomeDir() string {
	return userKit.GetUserHomeDir()
}

// ExpandTilde 拓展路径中的"~"（"~"（波浪号） => 当前用户主目录）.
/*
e.g.
("")		=> "", nil
("~") 		=> "/Users/richelieu", nil
("~/a/b/c")	=> "/Users/richelieu/a/b/c", nil
*/
func ExpandTilde(path string) (string, error) {
	return homedir.Expand(path)
}
