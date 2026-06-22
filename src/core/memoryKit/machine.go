package memoryKit

import (
	"context"

	"github.com/shirou/gopsutil/v4/mem"
)

// GetMachineMemoryStat 获取（当前瞬间的）服务器内存状态.
/*
PS:
(1) Total = Available + Used（？？？存疑，yozo有台Linux不符合）
(2) UsedPercent: 内存使用率 e.g.50.903940200805664
(3) Free和Available的区别:
	简单来说，Free内存是未被使用且处于空闲状态的内存，而Available内存则包括了已经被使用但可以释放的内存，例如缓存和缓冲区等.
	Available内存是一个 "估计值" ，表示在不使用交换空间的情况下可以使用多少内存。

mem.VirtualMemoryStat 结构体的字段:
(1) Total		总内存
(2) Available	可用内存（未被使用且处于空闲状态的内存 + 已经被使用但可以释放的内存，例如缓存和缓冲区等）
(3) Used		已使用内存
(4) UsedPercent	内存使用百分比
(5) Free		空闲状态的内存
*/
func GetMachineMemoryStat() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

func GetMachineMemoryStatWithContext(ctx context.Context) (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemoryWithContext(ctx)
}

// GetMachineAvailableMemory 获取（当前瞬间的）服务器可用内存.
/*
	PS: 很具有参考性.
*/
func GetMachineAvailableMemory() (uint64, error) {
	stat, err := GetMachineMemoryStat()
	if err != nil {
		return 0, err
	}
	return stat.Available, nil
}

// GetMachineUsedPercent 获取（当前瞬间的）服务器已使用内存百分比.
/*
	PS: 很具有参考性.

	计算公式:
		ret.Used = ret.Total - ret.Available
		ret.UsedPercent = float64(ret.Used) / float64(ret.Total) * 100.0
*/
func GetMachineUsedPercent() (float64, error) {
	stat, err := GetMachineMemoryStat()
	if err != nil {
		return 0, err
	}
	return stat.UsedPercent, nil
}
