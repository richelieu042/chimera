package mailKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

func AssertEmail(email string) error {
	if err := validateKit.Email(email); err != nil {
		return errorKit.NewfWithSkip(1, "[%s] email(%s) is invalid with error(%s)", funcKit.GetFuncName(1), email, err.Error())
	}
	return nil
}
