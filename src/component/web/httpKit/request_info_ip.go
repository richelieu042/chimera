package httpKit

import (
	"net"
	"net/http"
	"strings"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

var (
	// GetRequestPublicIp 获取http请求的 "公网ip".
	/*
		涉及的请求头:
		(1) "X-Forwarded-For"
		(2) "X-Real-Ip"
	*/
	GetRequestPublicIp func(req *http.Request) string = netutil.GetRequestPublicIp
)

// GetRemoteIP 获取客户端IP地址（客户端的远程IP地址）.
/*
PS: 参考 gin's Context.RemoteIP().

e.g.
当客户端通过代理服务器连接时，RemoteIP() 返回代理服务器的 IP 地址
*/
func GetRemoteIP(r *http.Request) string {
	//ctx := &gin.Context{
	//	Request: r,
	//}
	//return ctx.RemoteIP()

	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return ""
	}
	return ip
}

// GetClientIP
/*
PS: 参考 gin's Context.ClientIP().
*/
func GetClientIP(r *http.Request) string {
	ip := GetClientIPFromHeader(r)
	if strKit.IsEmpty(ip) {
		ip = GetRemoteIP(r)
	}
	return ip
}

// GetClientIPFromHeader
/*
PS: 参考 gin's Context.ClientIP().
*/
func GetClientIPFromHeader(r *http.Request) string {
	//ctx := &gin.Context{
	//	Request: r,
	//}
	//return ctx.ClientIP()

	for _, headerName := range RemoteIPHeaders {
		ip, valid := validateHeader(GetHeader(r.Header, headerName))
		if valid {
			return ip
		}
	}
	return ""
}

func validateHeader(header string) (clientIP string, valid bool) {
	if header == "" {
		return "", false
	}

	items := strings.Split(header, ",")
	for i := len(items) - 1; i >= 0; i-- {
		ipStr := strings.TrimSpace(items[i])
		ip := net.ParseIP(ipStr)
		if ip == nil {
			break
		}

		// X-Forwarded-For is appended by proxy
		// Check IPs in reverse order and stop when find untrusted proxy
		if i == 0 {
			return ipStr, true
		}
	}
	return "", false
}
