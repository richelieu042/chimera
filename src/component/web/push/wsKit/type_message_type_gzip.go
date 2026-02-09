package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu042/chimera/v3/src/compress/gzipKit"
)

// NewGzipMessageType
/*
PS: 此种情况下，必定使用 websocket.BinaryMessage（二进制数据）.

@param compressThreshold 压缩阈值，单位: byte
*/
func NewGzipMessageType(level, compressThreshold int) (*MessageType, error) {
	if err := gzipKit.AssertValidLevel(level); err != nil {
		return nil, err
	}

	return &MessageType{
		value: websocket.BinaryMessage,
		gzipConfig: &gzipKit.Config{
			Level:             level,
			CompressThreshold: compressThreshold,
		},
		brotliConfig: nil,
	}, nil
}
