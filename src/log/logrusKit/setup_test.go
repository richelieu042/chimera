package logrusKit

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestMustSetUp(t *testing.T) {
	MustSetUp(&Config{
		Level:       "debug",
		EnableQuote: true,
		PrintBasic:  true,
	})

	DisableQuote(nil)
	logrus.Info("1\n2\n3\n")

	EnableQuote(nil)
	logrus.Info("1\n2\n3\n")
}
