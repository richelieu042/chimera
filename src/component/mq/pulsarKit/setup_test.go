package pulsarKit

import (
	"fmt"
	"testing"

	"github.com/richelieu042/chimera/v3/src/config/viperKit"
	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
)

func TestSetUp(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		fmt.Printf("working dir: %s\n", wd)
	}

	type config struct {
		Pulsar *Config `json:"pulsar"`
	}

	path := "_chimera-lib/config.yaml"
	c := &config{}
	if _, err := viperKit.UnmarshalFromFile(path, nil, c); err != nil {
		panic(err)
	}
	MustSetUp(c.Pulsar, &VerifyConfig{
		Topic: "test",
		Print: true,
	})
}
