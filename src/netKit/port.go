package netKit

import (
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

const (
	DefaultHttpPort = 80

	DefaultHttpsPort = 443

	// MaxPort 65535 == 0xFFFF
	MaxPort = 0xFFFF
)

// IsPort 是否是有效的端口号？(0, 65535]
/*
参考:
(1) Java，hutool中的NetUtil.isValidPort()
(2) Linux端口分配: https://blog.csdn.net/zh2508/article/details/104888743

0 			不使用
1–1023 		系统保留,只能由root用户使用
1024—4999 	由客户端程序自由分配
5000—65535 	由服务器端程序自由分配（65535 = 2 ^ 16 - 1）

@param obj 	(1) 支持的类型: reflect.Value、int、uint、string...
			(2) 可以为nil
*/
func IsPort(obj interface{}) bool {
	return validateKit.Port(obj) == nil
}
