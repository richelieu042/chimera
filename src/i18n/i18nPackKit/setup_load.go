package i18nPackKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu042/chimera/v3/src/i18n/i18nKit"
)

func LoadMessageFile(path string) (*i18n.MessageFile, error) {
	if innerBundle == nil {
		return nil, NotSetupError
	}

	return i18nKit.LoadMessageFile(innerBundle, path)
}

func LoadMessageFileBytes(content []byte, path string) (*i18n.MessageFile, error) {
	if innerBundle == nil {
		return nil, NotSetupError
	}

	return i18nKit.LoadMessageFileBytes(innerBundle, content, path)
}
