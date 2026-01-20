package processKit

import (
	"context"
	"strconv"

	"github.com/richelieu-yang/chimera/v3/src/command/cmdKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

// GetProcessCount 获取: (瞬时的值)系统中所有进程的数量.
func GetProcessCount() (int, error) {
	data, err := cmdKit.RunCombinedly(context.TODO(), "sh", "-c", "ps auxw | wc -l")
	if err != nil {
		return 0, err
	}
	str := strKit.TrimSpace(string(data))

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetProcessThreadCount 获取: (瞬时的值)系统中所有进程及其线程的数量.
func GetProcessThreadCount() (int, error) {
	return 0, errorKit.Newf("not yet realized")
}
