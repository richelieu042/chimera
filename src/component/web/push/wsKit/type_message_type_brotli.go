package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu042/chimera/v3/src/compress/brotliKit"
)

// NewBrotliMessageType
/*
PS: 此种情况下，必定使用 websocket.BinaryMessage（二进制数据）.

@param compressThreshold 压缩阈值，单位: byte
*/
func NewBrotliMessageType(level, compressThreshold int) (*MessageType, error) {
	if err := brotliKit.AssertValidLevel(level); err != nil {
		return nil, err
	}

	return &MessageType{
		value:      websocket.BinaryMessage,
		gzipConfig: nil,
		brotliConfig: &brotliKit.Config{
			Level:             level,
			CompressThreshold: compressThreshold,
		},
	}, nil
}
