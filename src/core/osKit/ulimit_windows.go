package osKit

import "github.com/richelieu042/chimera/v3/src/core/error/errKit"

func GetUlimitInfo() (string, error) {
	return "", errKit.Newf("not yet realized")
}

func GetMaxOpenFiles() (int, error) {
	return 0, errKit.Newf("not yet realized")
}

func GetMaxProcessThreadCountByUser() (int, error) {
	return 0, errKit.Newf("not yet realized")
}

func GetCoreFileSize() (string, error) {
	return "", errKit.Newf("not yet realized")
}
