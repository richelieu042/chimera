package i18nKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
)

// GetMessage
/*
@param langs 可以不传，此时将使用 bundle 的默认语言
*/
func GetMessage(bundle *i18n.Bundle, messageId string, langs ...string) (string, error) {
	if err := interfaceKit.AssertNotNil(bundle, "bundle"); err != nil {
		return "", err
	}

	localizer := i18n.NewLocalizer(bundle, langs...)
	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageId,
	})
	if err != nil {
		return "", err
	}
	return message, nil
}
