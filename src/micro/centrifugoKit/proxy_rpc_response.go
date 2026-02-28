package centrifugoKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu042/chimera/v3/src/micro/centrifugoKit/proxyproto"
)

// PackToRpcResponse
/*
@param base64Flag 是否对 传参jsonData 进行base64编码？
*/
func PackToRpcResponse(jsonData []byte, base64Flag bool) *proxyproto.RPCResponse {
	resp := &proxyproto.RPCResponse{
		Result: &proxyproto.RPCResult{
			B64Data: "",
			Data:    nil,
		},
		Error:      nil,
		Disconnect: nil,
	}

	if base64Flag {
		resp.Result.B64Data = base64Kit.EncodeToString(jsonData)
	} else {
		resp.Result.Data = jsonData
	}
	return resp
}

// PackToRpcResponseWithCustomError 返回自定义错误.
/*
Return custom error
	https://centrifugal.dev/docs/server/proxy#return-custom-error

@param code 有效范围: [400, 1999]
*/
func PackToRpcResponseWithCustomError(code uint32, message string, temporaryArgs ...bool) (*proxyproto.RPCResponse, error) {
	if code < 400 || code > 1999 {
		return nil, errorKit.Newf("error code(%d) isn't in range[400, 1999]", code)
	}

	temporary := false
	if len(temporaryArgs) > 0 {
		temporary = temporaryArgs[0]
	}

	return &proxyproto.RPCResponse{
		Result: nil,
		Error: &proxyproto.Error{
			Code:      code,
			Message:   message,
			Temporary: temporary,
		},
		Disconnect: nil,
	}, nil
}

// PackToRpcResponseWithCustomDisconnect 返回自定义断开.
/*
Return custom disconnect
	https://centrifugal.dev/docs/server/proxy#return-custom-disconnect

@param code		有效范围: [4000, 4999]
 			  	(1) [4000,4499]: 给客户端一个重新连接的建议.
			  		(a) 客户端收到后会断开连接，但会重连;
			  		(b) 不会触发前端的 disconnected 事件;
			 		(c) Code 和 Reason 可以从前端的 connecting 事件中得知.
			  	(2) [4500,4999]: terminal codes，客户端接收到它后不会重新连接.
			  		(a) 客户端收到后会断开连接，且不会重连;
			  		(b) 会触发前端的 disconnected 事件;
			  		(c) Code 和 Reason 可以从前端的 disconnected 事件中得知.
@param reason 	请记住，由于WebSocket协议的限制和离心机内部协议的需要，你需要保持断开原因字符串不超过32个ASCII符号(即最大32字节)。
*/
func PackToRpcResponseWithCustomDisconnect(code uint32, reason string) (*proxyproto.RPCResponse, error) {
	if code < 4000 || code > 4999 {
		return nil, errorKit.Newf("disconnect code(%d) isn't in range[4000, 4999]", code)
	}

	length := len(reason)
	if length > 32 {
		return nil, errorKit.Newf("disconnect reason(length: %d, value: %s) is too long", length, reason)
	}

	return &proxyproto.RPCResponse{
		Result: nil,
		Error:  nil,
		Disconnect: &proxyproto.Disconnect{
			Code:   code,
			Reason: reason,
		},
	}, nil
}
