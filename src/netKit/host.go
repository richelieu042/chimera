package netKit

import (
	"fmt"
	"net"
	"strconv"

	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

var (
	// SplitHost
	/*
		e.g.
			fmt.Println(SplitHost("localhost:8080")) // "localhost" "8080" <nil>
			fmt.Println(SplitHost("127.0.0.1")) // "" "" address 127.0.0.1: missing port in address
			fmt.Println(SplitHost("localhost")) // "" "" address localhost: missing port in address
	*/
	SplitHost func(host string) (hostname, port string, err error) = net.SplitHostPort
)

// JoinToHost
/*
@param hostname (1) 支持: ipv4、ipv6
				(2) ipv6的话，最前和最后不要带上中括号

e.g.
	fmt.Println(netKit.JoinToHost("127.0.0.1", 80)) // 127.0.0.1:80
	fmt.Println(netKit.JoinToHost("", 8888))        // :8888

e.g.1 ipv6
	("::1", 8888) => "[::1]:8888"
*/
func JoinToHost(hostname string, port int) string {
	//return fmt.Sprintf("%s:%d", hostname, port)

	return net.JoinHostPort(hostname, strconv.Itoa(port))
}

func IsHostname(str string) bool {
	return validateKit.Hostname(str) == nil
}

func IsHost(str string) bool {
	return validateKit.Host(str) == nil
}

// PolyfillHosts
/*
PS:
(1) host包含 ip 和 port. e.g."127.0.0.1:80"

@param minCount 至少要有几个元素？

@return 第1个返回值: 可能是一个新的slice实例
*/
func PolyfillHosts(hosts []string, minCount int) ([]string, error) {
	if minCount <= 0 {
		minCount = 1
	}

	hosts = sliceKit.PolyfillStringSlice(hosts)
	//if err := sliceKit.AssertNotEmpty(hosts, "hosts"); err != nil {
	//	return nil, err
	//}
	tag := fmt.Sprintf("required,gte=%d,unique,dive,hostname_port", minCount)
	if err := validateKit.Var(hosts, tag); err != nil {
		return nil, errorKit.Wrapf(err, "hosts is invalid")
	}
	return hosts, nil
}
