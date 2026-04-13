// Package jsonRespKit
//
// Deprecated: Use i18nRespKit instead.
//
// This package is frozen and no new functionality will be added.
package jsonRespKit

import (
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
)

// MustSetUp 必须初始化.
/*
Deprecated: Use i18nRespKit instead.
*/
func MustSetUp(respProvider RespProvider, options ...Option) {
	if err := SetUp(respProvider, options...); err != nil {
		console.Fatalf("Fail to set up, error: %+v", err)
	}
}

// SetUp
/*
Deprecated: Use i18nRespKit instead.
*/
func SetUp(respProvider RespProvider, options ...Option) (err error) {
	defer func() {
		if err != nil {
			provider = nil
			msgMap = make(map[string]string)
		}
	}()

	err = interfaceKit.AssertNotNil(respProvider, "respProvider")
	if err != nil {
		return err
	}
	provider = respProvider

	opts := loadOptions(options...)

	for _, path := range opts.filePaths {
		if err = readFile(path); err != nil {
			return err
		}
	}
	for _, fd := range opts.fileDataSlice {
		if err = readFileData(fd); err != nil {
			return err
		}
	}
	return nil
}
