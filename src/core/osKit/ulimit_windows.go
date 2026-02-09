package osKit

import "github.com/richelieu042/chimera/v3/src/core/errorKit"

func GetUlimitInfo() (string, error) {
	return "", errorKit.Newf("not yet realized")
}

func GetMaxOpenFiles() (int, error) {
	return 0, errorKit.Newf("not yet realized")
}

func GetMaxProcessThreadCountByUser() (int, error) {
	return 0, errorKit.Newf("not yet realized")
}

func GetCoreFileSize() (string, error) {
	return "", errorKit.Newf("not yet realized")
}
