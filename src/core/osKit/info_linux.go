package osKit

import (
	"context"
	"strconv"

	"github.com/richelieu-yang/chimera/v3/src/command/cmdKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

// GetPidMax 获取: 系统的pid最大值（作为系统范围内 进程 和 线程 总数的限制）.
/*
PS:
(1) 大多数Linux上的默认值: 32768
	32位系统: 最大值为 32768
	64位系统: 任何小于等于 2^22（PID_MAX_LIMIT，约 400 万）的值

命令:
cat /proc/sys/kernel/pid_max
sysctl kernel.pid_max
*/
func GetPidMax() (int, error) {
	data, err := cmdKit.RunCombinedly(context.TODO(), "sh", "-c", "cat /proc/sys/kernel/pid_max")
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

// GetThreadsMax 获取: 系统的最大线程数.
/*
命令:
cat /proc/sys/kernel/threads-max
sysctl kernel.threads-max
*/
func GetThreadsMax() (int, error) {
	data, err := cmdKit.RunCombinedly(context.TODO(), "sh", "-c", "cat /proc/sys/kernel/threads-max")
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

// GetMaxMapCount 获取: 一个进程可以拥有的最大内存映射区域数量（间接限制了线程数）.
/*
PS: @return 间接限制了线程数，因为每个线程都需要一些内存映射区域.

命令:
cat /proc/sys/vm/max_map_count
sysctl vm.max_map_count
*/
func GetMaxMapCount() (int, error) {
	data, err := cmdKit.RunCombinedly(context.TODO(), "sh", "-c", "cat /proc/sys/vm/max_map_count")
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
