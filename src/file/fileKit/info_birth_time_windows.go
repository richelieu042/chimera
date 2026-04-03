package fileKit

import (
	"os"
	"syscall"
	"time"
)

// GetBirthTime 获取文件（或目录）的创建时间.
/*
@param info os.FileInfo（由 os.Stat 或 os.Lstat 获取）
@return 创建时间, true；若无法获取则 fallback 返回 ModTime, false
*/
func GetBirthTime(info os.FileInfo) (time.Time, bool) {
	stat, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		// fallback 到 ModTime
		return info.ModTime(), false
	}
	return time.Unix(0, stat.CreationTime.Nanoseconds()), true
}
