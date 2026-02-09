package wsKit

import (
	"context"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/richelieu042/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu042/chimera/v3/src/compress/brotliKit"
	"github.com/richelieu042/chimera/v3/src/compress/gzipKit"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
)

var (
	_ pushKit.Channel = &WsChannel{}
)

type WsChannel struct {
	pushKit.BaseChannel

	msgType *MessageType

	conn *websocket.Conn
}

func (channel *WsChannel) Initialize() error {
	if channel.PongInterval > 0 {
		channel.Interval = timeKit.SetInterval(context.TODO(), func(t time.Time) {
			if err := channel.Push(pushKit.PongData); err != nil {
				pushKit.GetDefaultLogger().WithError(err).Error("Fail to pong.")
				return
			}
		}, channel.PongInterval)
	}
	return nil
}

func (channel *WsChannel) Push(data []byte) error {
	return channel.PushMessage(channel.msgType, data)
}

// PushMessage 推送消息给客户端.
/*
@param MessageType MessageTypeText || MessageTypeBinary
*/
func (channel *WsChannel) PushMessage(messageType *MessageType, data []byte) (err error) {
	if messageType.gzipConfig != nil {
		/* gzip压缩 */
		data, err = gzipKit.Compress(data, gzipKit.WithLevel(messageType.gzipConfig.Level), gzipKit.WithCompressThreshold(messageType.gzipConfig.CompressThreshold))
		if err != nil {
			return err
		}
	} else if messageType.brotliConfig != nil {
		/* brotli压缩 */
		data, err = brotliKit.Compress(data, brotliKit.WithLevel(messageType.brotliConfig.Level), brotliKit.WithCompressThreshold(messageType.brotliConfig.CompressThreshold))
		if err != nil {
			return err
		}
	}

	if channel.Closed {
		return pushKit.ChannelClosedError
	}

	abortFlag := false

	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			err = pushKit.ChannelClosedError
			abortFlag = true
			return
		}

		err = channel.conn.WriteMessage(messageType.value, data)
	})

	if abortFlag {
		return
	}
	if err != nil {
		// Closed == true 的情况下，推送消息失败（基本上就是连接断开了）
		closeInfo := fmt.Sprintf("Fail to push because of error(%s)", err.Error())
		if channel.SetClosed() {
			channel.GetCloseCh() <- closeInfo
		}
	}
	return
}

// Close 后端主动关闭通道.
func (channel *WsChannel) Close(reason string) (err error) {
	closeInfo := fmt.Sprintf("Closed by backend with reason(%s)", reason)

	channel.Listeners.BeforeClosedByBackend(channel, closeInfo)

	// Richelieu: 这里延迟3毫秒，以免: BeforeClosedByBackend()里面会向前端发消息，前端先触发onclose，再触发onmessage
	time.Sleep(time.Millisecond * 3)

	if channel.SetClosed() {
		channel.GetCloseCh() <- closeInfo
		err = channel.conn.Close()
	}
	return
}

//func (channel *WsChannel) BindGroup(group string) {
//	pushKit.BindGroup(channel, group)
//}
//
//func (channel *WsChannel) BindUser(user string) {
//	pushKit.BindUser(channel, user)
//}
//
//func (channel *WsChannel) BindBsid(bsid string) {
//	pushKit.BindBsid(channel, bsid)
//}
