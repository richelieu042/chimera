//go:build 386 || amd64 || arm || arm64

package cpuKit

import (
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/load"
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

// GetLoadStat 获取操作系统的负载（1分钟、5分钟、15分钟）.
/*
负载值本身怎么理解？
	（1）Load 值代表处于可运行状态（R）或不可中断睡眠状态（D，通常是等待 I/O）的进程数量的指数衰减移动平均。
	（2）这个数值不是百分比，而是一个"需求量"：
		如果 Load1 = 1.0，大致意味着平均有 1 个进程在排队使用 CPU（或等 I/O）。
		判断负载是否"高"，需要结合 CPU 核心数来看：例如 4 核机器，Load1 = 4.0 大致相当于 CPU 被跑满；如果 Load1 = 8.0，说明已经严重过载，任务在排队等待。
	（3）一般经验法则：
		Load1 < 核心数：负载正常
		Load1 ≈ 核心数：CPU 接近饱和
		Load1 > 核心数：出现排队，系统响应可能变慢

三个时间窗口的作用
	Load1：反映最近的瞬时压力，波动较大，适合看突发情况。
	Load5：中期趋势，比 Load1 更平滑。
	Load15：长期趋势，用于判断是持续性高负载还是偶发尖峰。

常见判断技巧：
	如果 Load1 > Load5 > Load15：负载正在上升。
	如果 Load1 < Load5 < Load15：负载正在下降，可能是刚过了一个高峰。
*/
func GetLoadStat() (*load.AvgStat, error) {
	stat, err := load.Avg()
	if err != nil {
		return nil, err
	}
	return stat, nil
}

// GetLoadString
/*
@return e.g."0.12, 0.08, 0.08"
*/
func GetLoadString() (string, error) {
	stat, err := GetLoadStat()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f, %.2f, %.2f", stat.Load1, stat.Load5, stat.Load15), nil
}
