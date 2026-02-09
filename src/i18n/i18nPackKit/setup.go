package i18nPackKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/i18n/i18nKit"
	"golang.org/x/text/language"
)

var (
	NotSetupError = errorKit.Newf("haven’t been set up correctly")
)

var (
	innerBundle *i18n.Bundle

	innerBeanMaker BeanMaker

	defaultBeanMaker = func(code, msg string, data interface{}) interface{} {
		return &bean{
			Code:    code,
			Message: msg,
			Data:    data,
		}
	}
)

// SetUp
/*
@param maker 可以为nil
*/
func SetUp(defaultLanguage language.Tag, maker BeanMaker) {
	innerBundle = i18nKit.NewBundle(defaultLanguage)
	innerBeanMaker = maker
}

func getMaker() BeanMaker {
	if innerBeanMaker != nil {
		return innerBeanMaker
	}
	return defaultBeanMaker
}
