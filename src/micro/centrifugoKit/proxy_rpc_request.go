package centrifugoKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu042/chimera/v3/src/micro/centrifugoKit/proxyproto"
)

// ParseToRpcRequest
/*
一般情况下：
(1) centrifugo的 proxy_binary_encoding 配置为 false（默认），请求参数取 RPCRequest.Data
(2) centrifugo的 proxy_binary_encoding 配置为 true，请求参数取 RPCRequest.B64Data，还需要手动base64解码一下(with base64.StdEncoding)

@return err == nil的情况下，	获取 method 使用 rpcRequest.Method
							获取 data 使用rpcRequest.Data
*/
func ParseToRpcRequest(ctx *gin.Context) (*proxyproto.RPCRequest, error) {
	rpcRequest := &proxyproto.RPCRequest{}
	if err := ctx.Bind(rpcRequest); err != nil {
		return nil, err
	}
	// Richelieu: 此时，rpcRequest.Data 和 rpcRequest.B64Data 只有一个有值

	if rpcRequest.B64Data != "" {
		// Centrifugo服务的 proxy_binary_encoding 配置项为 true 的情况
		tmp, err := base64Kit.DecodeString(rpcRequest.B64Data)
		if err != nil {
			return nil, errorKit.Wrapf(err, "fail to decode B64Data")
		}
		rpcRequest.Data = tmp
	}
	return rpcRequest, nil
}
