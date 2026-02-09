package rocketmq5Kit

import (
	"fmt"
	"testing"

	"github.com/richelieu042/chimera/v3/src/config/viperKit"
	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
	"github.com/sirupsen/logrus"
)

func TestMustSetUp(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		logrus.Infof("wd: [%s].", wd)
	}

	path := "_chimera-lib/config.yaml"
	type config struct {
		RocketMQ5 *Config `json:"rocketmq5"`
	}
	c := &config{}
	_, err := viperKit.UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonKit.MarshalIndentToString(c.RocketMQ5, "", "    "))

	MustSetUp(c.RocketMQ5, "_client.log", &VerifyConfig{
		Topic:   "test",
		LogPath: "",
	})
}
