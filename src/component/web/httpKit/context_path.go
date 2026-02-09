package httpKit

import (
	"path"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// PolyfillContextPath
/**
 * e.g.
 * "" 				=>	""
 * "/"				=>	"/"
 * "//"				=>	"/"
 * "///"			=>	"/"
 * "/c/////c//"		=>	"/c/c/"
 *
 * @return 优化过的ContextPath
 */
func PolyfillContextPath(relativePath string) string {
	rst := joinPaths("", strKit.TrimSpace(relativePath))

	switch rst {
	case "":
	case "/":
	default:
		rst = strKit.PrependIfMissing(rst, "/")
		rst = strKit.AppendIfMissing(rst, "/")
	}
	return rst
}

// joinPaths
// 参考：github.com/gin-gonic/gin	routergroup.go	calculateAbsolutePath()
func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func lastChar(str string) byte {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}
