package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu042/chimera/v3/src/compress/brotliKit"
	"github.com/richelieu042/chimera/v3/src/compress/gzipKit"
)

type MessageType struct {
	value int

	gzipConfig   *gzipKit.Config
	brotliConfig *brotliKit.Config
}

func (msgType *MessageType) String() {

}

var (
	MessageTypeText = &MessageType{
		value:        websocket.TextMessage,
		gzipConfig:   nil,
		brotliConfig: nil,
	}

	MessageTypeBinary = &MessageType{
		value:        websocket.BinaryMessage,
		gzipConfig:   nil,
		brotliConfig: nil,
	}
)
