// Package runtimeKit 主要是对如下包的封装："runtime"
package runtimeKit

import (
	"runtime"

	"github.com/shirou/gopsutil/v4/host"
)

func GetHostInfo() (*host.InfoStat, error) {
	return host.Info()
}

// GetGoRoot 环境变量GOROOT
func GetGoRoot() string {
	return runtime.GOROOT()
}

// GetGoVersion Golang的版本号
func GetGoVersion() string {
	return runtime.Version()
}
