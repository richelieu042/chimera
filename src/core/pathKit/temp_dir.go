package pathKit

import (
	"os"

	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/idKit"
)

var tempDir string
var innerXid string = idKit.NewXid()

// GetOsTempDir 获取系统的临时目录.
/*
PS: 不建议向系统临时目录中放东西（服务器不一定会给权限）.

e.g.
	Windows	"C:\Users\Lenovo\AppData\Local\Temp"
	Mac		"/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/"
*/
var GetOsTempDir func() string = os.TempDir

// GetTempDir 获取 项目 的临时目录.
func GetTempDir() string {
	if strKit.IsNotEmpty(tempDir) {
		return tempDir
	}
	return GetOsTempDir()
}

func SetTempDir(dirPath string) error {
	if err := fileKit.AssertNotExistOrIsDir(dirPath); err != nil {
		return err
	}
	if err := fileKit.MkDirs(dirPath); err != nil {
		return err
	}

	tempDir = dirPath
	return nil
}

// GetExclusiveTempDir 获取 本依赖 的专属临时目录.
/*
	PS: 不建议向系统临时目录中放东西（服务器不一定会给权限）.

	e.g. macOS
	() => "/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/$$chimera/cmnmk1pus0n0snnatg40 <nil>", nil
*/
func GetExclusiveTempDir() (string, error) {
	dir := Join(GetTempDir(), "$$"+consts.ProjectName, innerXid)
	if err := fileKit.MkDirs(dir); err != nil {
		return "", err
	}
	return dir, nil
}
