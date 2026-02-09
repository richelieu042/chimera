package i18nKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"golang.org/x/text/language"
)

// Associate 关联其他语言代码.
/*
@param languageCode 语言代码，(1) 可参考: i18nKit/_info.md
							(2) 大小写“不敏感”
*/
func Associate(bundle *i18n.Bundle, messageFile *i18n.MessageFile, languageCodes ...string) (err error) {
	if err = interfaceKit.AssertNotNil(bundle, "bundle"); err != nil {
		return
	}
	if err = interfaceKit.AssertNotNil(messageFile, "messageFile"); err != nil {
		return
	}

	for i, languageCode := range languageCodes {
		if strKit.IsEmpty(languageCode) {
			err = errorKit.Newf("languageCode(index: %d, value: %s) is invalid", i, languageCode)
			return
		}
		tag := language.Make(languageCode)
		if err = bundle.AddMessages(tag, messageFile.Messages...); err != nil {
			err = errorKit.Wrapf(err, "AddMessages() fails with languageCode(index: %d, value: %s)", i, languageCode)
			return
		}
	}
	return
}
