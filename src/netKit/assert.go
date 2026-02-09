package netKit

import (
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

func AssertValidPort(port int) error {
	if !IsPort(int64(port)) {
		return errorKit.NewfWithSkip(1, "[%s] port(%d) is invalid", funcKit.GetFuncName(1), port)
	}
	return nil
}

func AssertHost(host string) error {
	if !IsHost(host) {
		return errorKit.NewfWithSkip(1, "[%s] host(%s) is invalid", funcKit.GetFuncName(1), host)
	}
	return nil
}

func AssertHostname(hostname string) error {
	if !IsHostname(hostname) {
		return errorKit.NewfWithSkip(1, "[%s] hostname(%s) is invalid", funcKit.GetFuncName(1), hostname)
	}
	return nil
}
