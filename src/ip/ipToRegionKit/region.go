package ipToRegionKit

// GetIPv4Region
/*
PS:
(1) 需要先加载 ip2region.xdb
(2) 此函数可以安全用于并发
(3) IP地址查询 - 在线工具（可以查看本机的外网ip）: https://tool.lu/ip/

@return 第1个返回值的格式: 国家|区域|省份|城市|ISP

e.g.
("1.1.1.1")			=> ("澳大利亚|0|0|0|0", nil)
("1.2.3.4") 		=> ("美国|0|华盛顿|0|谷歌", nil)
("10.0.9.141") 		=> ("0|0|0|内网IP|内网IP", nil)
("218.90.174.146") 	=> ("中国|0|江苏省|无锡市|电信", nil)
*/
func GetIPv4Region(ipv4 string) (string, error) {
	if ipv4Searcher == nil {
		return "", NotSetupError
	}

	// 每个 ipv4 数据段的 region 信息都固定了格式：国家|区域|省份|城市|ISP，只有中国的数据绝大部分精确到了城市，其他国家部分数据只能定位到国家，后前的选项全部是0。
	return ipv4Searcher.SearchByStr(ipv4)
}

func GetIPv6Region(ipv6 string) (string, error) {
	if ipv4Searcher == nil {
		return "", NotSetupError
	}

	// 每个 ipv6 数据段的 region 信息都固定了格式：国家|区域|省份|城市|ISP，只有中国的数据绝大部分精确到了城市，其他国家部分数据只能定位到国家，后前的选项全部是0。
	return ipv6Searcher.SearchByStr(ipv6)
}
