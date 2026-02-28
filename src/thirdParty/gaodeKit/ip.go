package gaodeKit

import (
	"context"

	"github.com/richelieu042/chimera/v3/src/component/web/http_client/reqKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/ip/ipKit"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonFieldKit"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
	"github.com/tidwall/gjson"
)

const (
	ipUrl = "https://restapi.amap.com/v3/ip"
)

// GetIpInfo
/*
PS: 仅支持IPV4，不支持国外IP解析.
*/
func (client *Client) GetIpInfo(ip string) (*IpInfo, error) {
	if err := ipKit.AssertIPv4(ip); err != nil {
		return nil, err
	}

	resp := reqKit.Get(context.TODO(), ipUrl, map[string][]string{
		"key": {client.key},
		"ip":  {ip},
	})
	if resp.Err != nil {
		return nil, resp.Err
	}
	jsonData, err := resp.ToBytes()
	if err != nil {
		return nil, err
	}

	/*
		局域网ip响应:	{"status":"1","info":"OK","infocode":"10000","province":"局域网","city":[],"adcode":[],"rectangle":[]}
		外网ip响应: 	{"status":"1","info":"OK","infocode":"10000","province":[],"city":[],"adcode":[],"rectangle":[]}
	*/
	field := jsonFieldKit.GetField(jsonData, "province")
	if field.Type == gjson.String {
		if field.String() == "局域网" {
			return &IpInfo{
				Province:  "局域网",
				City:      "",
				Adcode:    "",
				Rectangle: "",
			}, nil
		} else {
			// continue
		}
	} else {
		if field.IsArray() && field.Raw == "[]" {
			return &IpInfo{
				Province:  "外网",
				City:      "",
				Adcode:    "",
				Rectangle: "",
			}, nil
		}
		return nil, errorKit.Newf("invalid response(%s)", string(jsonData))
	}

	/* 国内ip */
	ipResp := &IpResponse{}
	if err := jsonKit.Unmarshal(jsonData, ipResp); err != nil {
		return nil, errorKit.Wrapf(err, "Fail to unmarshal")
	}
	if err := ipResp.IsSuccess(); err != nil {
		return nil, err
	}
	return &ipResp.IpInfo, nil
}
