package ip2RegionKit

import (
	"testing"

	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
)

func TestGetRegion(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		console.Infof("working dir: [%s].", wd)
	}

	/*
		https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.xdb
	*/
	ipv4XdbPath := "_temp/ip2region_v4.xdb"
	ipv6XdbPath := "_temp/ip2region_v6.xdb"
	MustSetUp(ipv4XdbPath, ipv6XdbPath)

	{
		ip := "155.117.18.75"
		str, err := GetIPv4Region(ip)
		if err != nil {
			panic(err)
		}
		console.Info(str)
	}
	{
		ip := "2602:f988:210:251::c30"
		str, err := GetIPv6Region(ip)
		if err != nil {
			panic(err)
		}
		console.Info(str)
	}
}
