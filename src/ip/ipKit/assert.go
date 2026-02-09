package ipKit

import (
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

func AssertIP(ip string) error {
	if err := validateKit.IP(ip); err != nil {
		return errorKit.NewfWithSkip(1, "[%s] ip(%s) is invalid because of error(%s)", funcKit.GetFuncName(1), ip, err.Error())
	}
	return nil
}

func AssertIPv4(ipv4 string) error {
	if err := validateKit.IPv4(ipv4); err != nil {
		return errorKit.NewfWithSkip(1, "[%s] ipv4(%s) is invalid because of error(%s)", funcKit.GetFuncName(1), ipv4, err.Error())
	}
	return nil
}
