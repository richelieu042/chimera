package ipKit

import (
	"github.com/duke-git/lancet/v2/netutil"
)

var (
	// GetIps 获取ipv4地址列表.
	/*
		e.g.
		fmt.Println(ipKit.GetIps()) // [172.20.10.4 198.18.0.1]
	*/
	GetIps func() []string = netutil.GetIps
)

//// GetLocalIPs
///*
//流程:
//	使用 net.InterfaceAddrs() 来获取所有网卡的地址，然后遍历这些地址，找到 IPv4地址 且 不是回环地址 的IP地址
//缺陷:
//	这种方法无法直接获取到对外的IP地址，只能判断IPv4和非回环地址，多IP情况还需要额外进行判断。
//*/
//func GetLocalIPs() ([]string, error) {
//	addrs, err := net.InterfaceAddrs()
//	if err != nil {
//		return nil, err
//	}
//
//	var ips []string
//	for _, address := range addrs {
//		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
//			if ipNet.IP.To4() != nil {
//				ips = append(ips, ipNet.IP.String())
//			}
//		}
//	}
//	if ips == nil {
//		return nil, errKit.Newf("fail to get local ips")
//	}
//	return ips, nil
//}
//
//var outboundIp string
//var outboundIpErr error
//var outboundOnce sync.Once
//
//// GetOutboundIP 获取: 对外的ip地址
///*
//系统性能数据gopsutil库
//	https://topgoer.com/%E5%85%B6%E4%BB%96/%E7%B3%BB%E7%BB%9F%E6%80%A7%E8%83%BD%E6%95%B0%E6%8D%AEgopsutil%E5%BA%93.html
//通过 UDP 获取本机 IP
//	http://t.zoukankan.com/fousor-p-14874576.html
//
//PS: 由于会申请一个UDP的端口，所以如果经常调用也会比较耗时的，这里如果需要可以将查询到的IP给缓存起来，性能可以获得很大提升.
//
//流程:
//(1) 使用 net.Dial 连接到一个外部地址（例如“8.8.8.8:53”），
//(2) 然后通过 conn.LocalAddr() 获取到本地IP地址.
//
//优点:
//	这种方法可以直接获取到对外的IP地址.
//	(a) 使用 UDP 的优点是不需要关注是否送达，只需要对应的 {ip}:{port} 结构正确，即可获取到IP地址；
//	(b) 这里使用TCP 也是可以的，只是需要保证对应的 {ip}:{port} 连通性.
//*/
//func GetOutboundIP() (string, error) {
//	outboundOnce.Do(func() {
//		callback := func() (string, error) {
//			addr := "8.8.8.8:80"
//
//			// 使用UDP（不需要保证addr的连通性）
//			conn, err := net.Dial("udp", addr)
//			if err != nil {
//				return "", err
//			}
//			defer conn.Close()
//			localAddr := conn.LocalAddr().(*net.UDPAddr)
//			return localAddr.IP.String(), nil
//		}
//		outboundIp, outboundIpErr = callback()
//	})
//
//	return outboundIp, outboundIpErr
//}
