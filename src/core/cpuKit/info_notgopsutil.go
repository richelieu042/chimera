//go:build !(386 || amd64 || arm || arm64)

package cpuKit

import (
	"os"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
)

// GetUsagePercent CPU使用率
func GetUsagePercent() (float64, error) {
	return 0, errorKit.Newf("Not supported on this platform")
}

// GetCurrentProcessUsagePercent 获取 当前进程 的CPU使用百分比.
/*
PS: 类似Linux命令: top -p ${pid}
*/
func GetCurrentProcessUsagePercent() (float64, error) {
	var pid = int32(os.Getpid())
	return GetProcessUsagePercent(pid)
}

// GetProcessUsagePercent 获取 指定进程 的CPU使用百分比.
func GetProcessUsagePercent(pid int32) (float64, error) {
	return 0, errorKit.Newf("Not supported on this platform")
}
