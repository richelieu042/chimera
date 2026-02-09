package ipKit

import (
	"net"

	"github.com/richelieu042/chimera/v3/src/ip/ipKit/ipType"
)

type (
	IpInfo struct {
		Type int
		IP   net.IP
	}
)

func (info IpInfo) IsIP() bool {
	return info.IsIPv4() || info.IsIPv6()
}

func (info IpInfo) IsIPv4() bool {
	return info.Type == ipType.IPv4
}

func (info IpInfo) IsIPv6() bool {
	return info.Type == ipType.IPv6
}
