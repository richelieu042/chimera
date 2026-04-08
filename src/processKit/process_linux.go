package processKit

import (
	"context"
	"strconv"

	"github.com/richelieu042/chimera/v3/src/command/cmdKit"
)

// GetProcessCount 获取: (瞬时的值)系统中所有进程的数量.
/*
支持: 	Linux、Mac
*/
func GetProcessCount() (int, error) {
	str, _, err := cmdKit.RunToString(context.TODO(), "sh", "-c", "ps auxw | wc -l")
	if err != nil {
		return 0, err
	}

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
	str, _, err := cmdKit.RunToString(context.TODO(), "sh", "-c", "ps -eLf | wc -l")
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}
