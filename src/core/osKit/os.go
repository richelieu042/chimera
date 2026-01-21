package osKit

import (
	"runtime"

	"github.com/duke-git/lancet/v2/system"
)

// OS 操作系统
var OS string = runtime.GOOS

// ARCH 处理器架构
var ARCH string = runtime.GOARCH

var (
	// IsWindows 检查当前操作系统是否是windows
	IsWindows func() bool = system.IsWindows

	// IsLinux 检查当前操作系统是否是linux.
	IsLinux func() bool = system.IsLinux

	// IsMac 检查当前操作系统是否是macos.
	IsMac func() bool = system.IsMac
)

// GetOsBits 获取当前操作系统位数，返回32或64.
/*
	@return 32 || 64
*/
func GetOsBits() int {
	//return strconv.IntSize
	return system.GetOsBits()
}
