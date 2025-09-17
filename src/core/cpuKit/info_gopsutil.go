//go:build 386 || amd64 || arm || arm64

package cpuKit

import (
	"os"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/process"
)

// GetUsagePercent CPU使用率
/*
PS: 耗时约1s.

e.g.
() => 12.701612903175233
*/
func GetUsagePercent() (float64, error) {
	s, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return s[0], nil
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
	p, err := process.NewProcess(pid)
	if err != nil {
		return 0, err
	}
	return p.CPUPercent()
}
