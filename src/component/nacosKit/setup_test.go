package nacosKit

import (
	"testing"

	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/richelieu042/chimera/v3/src/config/viperKit"
	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
	"github.com/richelieu042/chimera/v3/src/log/logrusKit"
	"github.com/richelieu042/chimera/v3/src/randomKit"
	"github.com/sirupsen/logrus"
)

func TestMustSetUp(t *testing.T) {
	logrusKit.MustSetUp(nil)

	wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}
	logrus.Infof("wd: [%s].", wd)
	path := "_chimera-lib/config.yaml"

	type config struct {
		Nacos *Config `json:"nacos"`
	}

	c := &config{}
	if _, err := viperKit.UnmarshalFromFile(path, nil, c); err != nil {
		panic(err)
	}
	MustSetUp(c.Nacos)

	namingClient, err := NewNamingClient()
	if err != nil {
		logrus.Fatal(err)
	}
	defer namingClient.CloseClient()
	configClient, err := NewConfigClient()
	if err != nil {
		logrus.Fatal(err)
	}
	defer configClient.CloseClient()

	//instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
	//	ServiceName: "ws",
	//	GroupName:   "wo3",
	//})
	//fmt.Println(instance.Weight)
	//fmt.Println(instance.Metadata)

	flag, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:   "127.0.0.1",
		Port: 12345,
		/*
			有效范围: [0.0, 10000, 0]
		*/
		Weight:   float64(randomKit.Int(0, 1000)),
		Enable:   true,
		Healthy:  true,
		Metadata: nil,
		//ClusterName: "",
		ServiceName: "s",
		GroupName:   "g",
		Ephemeral:   false,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("flag: [%t].", flag)
}
