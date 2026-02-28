//go:build !(386 || amd64 || arm || arm64)

package diskKit

import (
	"github.com/richelieu042/chimera/v3/src/core/conditionKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/osKit"
)

func GetDiskUsageStats() (*DiskUsageStats, error) {
	path := conditionKit.TernaryOperator(osKit.IsWindows(), "C:", "/")
	return GetDiskUsageStatsByPath(path)
}

func GetDiskUsageStatsByPath(path string) (*DiskUsageStats, error) {
	return nil, errorKit.Newf("Currently not supported")
}
