package fileKit

import (
	"os"
	"time"

	"github.com/gogf/gf/v2/os/gfile"
)

var (
	// Chmod 使用指定的权限，更改指定路径的文件权限.
	Chmod func(path string, mode os.FileMode) (err error) = gfile.Chmod

	// CutAndPaste 剪贴.
	CutAndPaste func(src string, dst string) (err error) = gfile.Move

	// Truncate 更改文件大小的函数.
	/*
		其主要作用是：
		(1) 截断文件：将文件大小缩小到指定的长度，如果文件的当前内容长度超过指定长度，多余的内容会被直接截掉。
		(2) 扩展文件：如果指定的长度大于当前文件长度，文件将被扩展，新增的部分会用空字节（即 \x00）填充。

		PS:
		(1) 如果给定文件路径是软链，将会修改源文件;
		(2) If there is an error, it will be of type *PathError.

		@param size 如果为0，则清空文件内容
	*/
	Truncate func(path string, size int) (err error) = gfile.Truncate
)

// SetModificationTime 修改文件（或目录）的修改时间
/*
PS:
(1) 也会同时修改文件（或目录）的访问时间；
(2) 修改目录的修改时间，将不会影响该目录下的文件或目录；
(3) 传参t可以晚于当前时间.

@param path 传参""将返回error（chtimes : The system cannot find the path specified.）
*/
func SetModificationTime(path string, t time.Time) error {
	return os.Chtimes(path, t, t)
}
