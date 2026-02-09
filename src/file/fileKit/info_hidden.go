//go:build !windows

package fileKit

import (
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// IsHidden 文件（或目录）是否隐藏？
/*
如何在 Go 中检测文件夹中的隐藏文件 - 跨平台方法
	https://www.likecs.com/ask-919454.html#sc=1368.5

流程:
(1) 获取文件名（以防传参为路径）
(2) 判断文件名是否以"."开头

PS:
(1) 传参path 对应的文件或目录必须存在，否则返回error.

@param 文件（或目录）的 path 或 name
*/
func IsHidden(path string) (bool, error) {
	if err := AssertExist(path); err != nil {
		return false, err
	}

	return strKit.StartWith(GetFileName(path), "."), nil
}
