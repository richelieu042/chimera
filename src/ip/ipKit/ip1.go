package ipKit

//import (
//	"github.com/richelieu042/chimera/v3/src/ip/ipKit/ipType"
//	"net"
//)
//
//// GetIpInfo 获取ip字符串的信息.
///*
//@param str 字符串，IPv4 或 IPv6
//*/
//func GetIpInfo(str string) IpInfo {
//	kind := ipType.UNKNOWN
//	ip := net.ParseIP(str)
//	if ip != nil {
//		// ！！！：加入标签才能break 外面的循环
//		//OuterLoop:
//		//	for _, c := range str {
//		//		switch c {
//		//		case '.':
//		//			kind = ipType.IPv4
//		//			break OuterLoop
//		//		case ':':
//		//			kind = ipType.IPv6
//		//			break OuterLoop
//		//		default:
//		//			break
//		//		}
//		//	}
//
//		// If ip is not an IPv4 address, To4 returns nil.
//		if ip.To4() != nil {
//			kind = ipType.IPv4
//		} else {
//			kind = ipType.IPv6
//		}
//	}
//
//	return IpInfo{
//		Type: kind,
//		IP:   ip,
//	}
//}
//
//// IsIP 检查传参是否为ip（支持ipv4和ipv6）？
///*
//参考：https://www.php.cn/be/go/441401.html
//*/
//func IsIP(str string) bool {
//	return net.ParseIP(str) != nil
//}
//
//// IsIPv4
///*
//"0.0.0.0" 	=> true
//"127.0.0.1"	=> true
//*/
//func IsIPv4(str string) bool {
//	ip := net.ParseIP(str)
//	return ip != nil && ip.To4() != nil
//}
//
//func IsIPv6(str string) bool {
//	ip := net.ParseIP(str)
//	return ip != nil && ip.To4() == nil
//}
