package httpKit

import (
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

func AssertHttpUrl(httpUrl string) error {
	if err := validateKit.HttpUrl(httpUrl); err != nil {
		return errorKit.NewfWithSkip(1, "[%s] httpUrl(%s) is because of with error(%s)", funcKit.GetFuncName(1), httpUrl, err.Error())
	}
	return nil
}
