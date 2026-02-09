package jsonKit

import (
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

func AssertJson(data []byte, paramName string) error {
	if !IsJson(data) {
		return errorKit.NewfWithSkip(1, "[%s] param(name: %s) isn't a json, value: %s",
			funcKit.GetFuncName(1), paramName, string(data))
	}
	return nil
}

func AssertJsonString(jsonStr string, paramName string) error {
	if !IsJsonString(jsonStr) {
		return errorKit.NewfWithSkip(1, "[%s] param(name: %s) isn't a json, value: %s",
			funcKit.GetFuncName(1), paramName, jsonStr)
	}
	return nil
}
