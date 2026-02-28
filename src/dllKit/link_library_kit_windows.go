package linkLibraryKit

import (
	"plugin"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
)

// LoadLinkLibrary 加载动态链接库（Linux、Mac）
/*
TODO: 看后续"plugin标准库"是否会支持Windows环境.
*/
func LoadLinkLibrary(path string) (*plugin.Plugin, error) {
	return nil, errorKit.Newf("Link libraries cannot be loaded in Windows!")
}
