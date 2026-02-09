package gaodeKit

import (
	"fmt"
	"testing"

	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

func TestClient_GetIp(t *testing.T) {
	client, err := NewClient("b15c36bf1df4c272e92f3f1875a127f1")
	if err != nil {
		panic(err)
	}

	/* 局域网ip */
	{
		ip := "10.0.9.141"
		ipInfo, err := client.GetIpInfo(ip)
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonKit.MarshalIndentToString(ipInfo, "", "    "))
	}
	/* 国内ip */
	{
		ip := "218.90.174.146"
		ipInfo, err := client.GetIpInfo(ip)
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonKit.MarshalIndentToString(ipInfo, "", "    "))
	}
	/* 日本ip */
	{
		ip := "1.0.16.0"
		ipInfo, err := client.GetIpInfo(ip)
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonKit.MarshalIndentToString(ipInfo, "", "    "))
	}
	/*
		{
		    "province": "局域网",
		    "city": "",
		    "adcode": "",
		    "rectangle": ""
		} <nil>
		{
		    "province": "江苏省",
		    "city": "无锡市",
		    "adcode": "320200",
		    "rectangle": "120.1788533,31.4648817;120.4605818,31.68307651"
		} <nil>
		{
		    "province": "外网",
		    "city": "",
		    "adcode": "",
		    "rectangle": ""
		} <nil>
	*/
}
