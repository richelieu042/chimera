//go:build 386 || amd64 || arm || arm64

package diskKit

import (
	"github.com/richelieu042/chimera/v3/src/core/conditionKit"
	"github.com/richelieu042/chimera/v3/src/core/osKit"
	"github.com/shirou/gopsutil/v4/disk"
)

func GetDiskUsageStats() (*DiskUsageStats, error) {
	path := conditionKit.TernaryOperator(osKit.IsWindows(), "C:", "/")
	return GetDiskUsageStatsByPath(path)
}

// GetDiskUsageStatsByPath
/*
PS:
(1) Mac（Linux），查看磁盘空间的命令: df -h

golang 获取cpu 内存 硬盘 使用率 信息 进程信息
	https://blog.csdn.net/whatday/article/details/109620192
*/
func GetDiskUsageStatsByPath(path string) (*DiskUsageStats, error) {
	usageStat, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	return &DiskUsageStats{
		Path:        path,
		Free:        usageStat.Free,
		Used:        usageStat.Used,
		Total:       usageStat.Total,
		UsedPercent: usageStat.UsedPercent,
	}, nil

	//parts, err := disk.Partitions(true)
	//if err != nil {
	//	return nil, err
	//}
	//for _, part := range parts {
	//	if part.Mountpoint != "/" {
	//		continue
	//	}
	//	usageStat, err := disk.Usage(part.Mountpoint)
	//	if err != nil {te
	//		return nil, err
	//	}
	//	return (*DiskUsageStats)(usageStat), nil
	//}
	//return nil, errorKit.Newf("fail to get disk usageStat with parts(length: %d)", len(parts))
}
