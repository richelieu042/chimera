package i18nPackKit

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/i18n/i18nKit"
)

// Pack
/*
@param langs 可以为空（包括nil）
*/
func Pack(langs []string, code string, data interface{}, msgArgs ...interface{}) interface{} {
	return PackFully(langs, code, "", data, msgArgs...)
}

// PackFully
/*
@param langs 可以为空（包括nil）
*/
func PackFully(langs []string, code, msg string, data interface{}, msgArgs ...interface{}) interface{} {
	var err error

	if strKit.IsEmpty(msg) {
		msg, err = i18nKit.GetMessage(innerBundle, code, langs...)
	}
	if err != nil {
		msg = err.Error()
	} else {
		if strKit.IsNotEmpty(msg) && msgArgs != nil {
			msg = fmt.Sprintf(msg, msgArgs...)
		}
	}
	return getMaker()(code, msg, data)
}
