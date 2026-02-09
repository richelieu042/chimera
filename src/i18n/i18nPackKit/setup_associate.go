package i18nPackKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu042/chimera/v3/src/i18n/i18nKit"
)

// Associate 关联其他语言代码.
/*
@param languageCode 语言代码，(1) 可参考: i18nKit/_info.md
							(2) 大小写“不敏感”
*/
func Associate(messageFile *i18n.MessageFile, languageCodes ...string) error {
	if innerBundle == nil {
		return NotSetupError
	}

	return i18nKit.Associate(innerBundle, messageFile, languageCodes...)
}
