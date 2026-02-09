package envKit

import (
	"io"

	"github.com/joho/godotenv"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

var (
	// Load 从 文件 中读取配置，并存储到程序的环境变量中.
	/*
		PS:
		(1) 默认情况下，加载的是项目根目录下的.env文件;
		(2) 如果多个文件中存在同一个键，那么先出现的优先，后出现的不生效;
		(3) 会存储到程序的环境变量中.
	*/
	Load func(filenames ...string) (err error) = godotenv.Load

	// Overload 类似于 Load，但是会覆盖 先前文件中已存在 的环境变量.
	Overload func(filenames ...string) (err error) = godotenv.Overload

	// ReadFromString 从 string 中读取配置.
	/*
		PS: "不会"存储到程序的环境变量中.
	*/
	ReadFromString func(str string) (envMap map[string]string, err error) = godotenv.Unmarshal

	// ParReadFromReader 从 io.Reader 中读取配置.
	/*
		PS: "不会"存储到程序的环境变量中.
	*/
	ParReadFromReader func(r io.Reader) (map[string]string, error) = godotenv.Parse
)

// ReadFromFile 从 文件 中读取配置.
/*
	PS: "不会"存储到程序的环境变量中.
*/
func ReadFromFile(paths ...string) (envMap map[string]string, err error) {
	for _, path := range paths {
		if err := fileKit.AssertExistAndIsFile(path); err != nil {
			return nil, err
		}
	}

	return godotenv.Read(paths...)
}
