package rsaKit

import (
	"testing"

	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
)

func TestGenerateKeyFiles(t *testing.T) {
	_, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}

	bits := 2048
	format := PKCS1
	password := "dqwdqwd强无敌群多"

	priPath := "_pri.pem"
	pubPath := "_pub.pem"
	if err := GenerateKeyFiles(bits, format, password, priPath, pubPath, 0644); err != nil {
		panic(err)
	}
}
