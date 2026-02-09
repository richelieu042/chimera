package linkLibraryKit

import (
	"plugin"

	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// LoadLinkLibrary 加载动态链接库（Linux、Mac）
func LoadLinkLibrary(path string) (*plugin.Plugin, error) {
	if strKit.IsEmpty(path) {
		return nil, errorKit.Newf("path of link library is empty")
	}
	if !fileKit.Exists(path) {
		return nil, errorKit.Newf("link library(path: %s) doesn't exist", path)
	}
	return plugin.Open(path)
}
