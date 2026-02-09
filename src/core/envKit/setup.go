package envKit

import (
	"github.com/richelieu042/chimera/v3/src/log/console"
)

func MustSetUp(envFilePaths ...string) {
	err := SetUp(envFilePaths...)
	if err != nil {
		console.Fatalf("Fail to set up, error: %s", err.Error())
	}
}

// SetUp
/*
PS:
(1) 默认情况下，加载的是项目根目录下的.env文件;
(2) 如果多个文件中存在同一个键，那么先出现的优先，后出现的不生效;
(3) 会存储到程序的环境变量中.

@params envFilePaths 可以为nil || []string{}，将采用默认值: []string{".env"}
*/
func SetUp(envFilePaths ...string) error {
	return Load(envFilePaths...)
}
