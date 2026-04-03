package fileKit

import (
	"os"
	"time"

	"golang.org/x/sys/unix"
)

// GetBirthTime 获取文件（或目录）的创建时间.
/*
Linux 早期 stat 系统调用的结构体设计时就没有 BirthTime，后来虽然在内核 4.11 加了新的 statx 系统调用来补充这个字段，
但 os.FileInfo.Sys() 返回的仍是老的 *syscall.Stat_t，信息不在里面。
*/
func GetBirthTime(info os.FileInfo) (time.Time, bool) {
	// 无法从 Stat_t 取创建时间，直接 fallback 到 ModTime
	return info.ModTime(), false
}

// GetBirthTimePath 获取文件（或目录）的创建时间.
/*
调用 statx，ext4/xfs 返回真实创建时间.

@return 创建时间, true；若内核或文件系统不支持则 fallback 返回 ModTime, false
*/
func GetBirthTimeWithPath(path string) (time.Time, bool) {
	var sx unix.Statx_t

	err := unix.Statx(unix.AT_FDCWD, path, unix.AT_STATX_SYNC_AS_STAT, unix.STATX_BTIME, &sx)
	if err != nil || sx.Mask&unix.STATX_BTIME == 0 || (sx.Btime.Sec == 0 && sx.Btime.Nsec == 0) {
		// fallback：重新 Stat 拿 ModTime
		info, statErr := os.Stat(path)
		if statErr != nil {
			return time.Time{}, false
		}
		return info.ModTime(), false
	}
	return time.Unix(int64(sx.Btime.Sec), int64(sx.Btime.Nsec)), true
}
