package processKit

import (
	"context"
	"strconv"

	"github.com/richelieu-yang/chimera/v3/src/command/cmdKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

// GetProcessCount 获取: (瞬时的值)系统中所有进程的数量.
/*
支持: 	Linux、Mac
*/
func GetProcessCount() (int, error) {
	str, err := cmdKit.RunToString(context.TODO(), "sh", "-c", "ps auxw | wc -l")
	if err != nil {
		return 0, err
	}
	str = strKit.TrimSpace(str)

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetProcessThreadCount 获取: (瞬时的值)系统中所有进程及其线程的数量.
/*
支持: 	Linux
不支持:	Mac
*/
func GetProcessThreadCount() (int, error) {
	str, err := cmdKit.RunToString(context.TODO(), "sh", "-c", "ps -eLf | wc -l")
	if err != nil {
		return 0, err
	}
	str = strKit.TrimSpace(str)

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}
