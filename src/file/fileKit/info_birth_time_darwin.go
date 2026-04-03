package fileKit

import (
	"os"
	"syscall"
	"time"
)

// GetBirthTime 获取文件（或目录）的创建时间.
/*
@param info os.FileInfo（由 os.Stat 或 os.Lstat 获取）
@return bool 是否成功获取到？
*/
func GetBirthTime(info os.FileInfo) (time.Time, bool) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		// fallback 到 ModTime
		return info.ModTime(), false
	}
	return time.Unix(stat.Birthtimespec.Sec, stat.Birthtimespec.Nsec), true
}
