package sseKit

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/component/web/ginKit"
	"github.com/richelieu042/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu042/chimera/v3/src/concurrency/poolKit"
	"github.com/richelieu042/chimera/v3/src/core/bytesKit"
	"github.com/sirupsen/logrus"
)

func TestNewProcessor(t *testing.T) {
	engine := gin.Default()
	engine.Use(ginKit.NewCorsMiddleware(nil, true))

	/* 初始化poolKit */
	pool, err := poolKit.NewAntPool(1024)
	if err != nil {
		panic(err)
	}
	pushKit.MustSetUp(pool, nil)

	/* SSE */
	//msgType := MessageTypeRaw
	msgType := MessageTypeBase64

	processor, err := NewProcessor(nil, &demoListener{}, msgType, time.Second*10)
	if err != nil {
		logrus.Fatal(err)
	}
	engine.GET("/sse", processor.ProcessWithGin)

	if err := engine.Run(":12000"); err != nil {
		logrus.Fatal(err)
	}
}

type demoListener struct {
	pushKit.Listener
}

func (l *demoListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	logrus.WithField("failureInfo", failureInfo).Error("OnFailure")
}

func (l *demoListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel pushKit.Channel) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),
	}).Info("OnHandshake")

	if err := channel.Push([]byte("test 测试")); err != nil {
		logrus.WithError(err).Error("Fail to push when on handshake.")
	}
}

func (l *demoListener) OnMessage(channel pushKit.Channel, messageType int, data []byte) {
	if bytesKit.Equals(data, pushKit.PingData) {
		return
	}

	msgText := string(data)
	if msgText == "close" {
		_ = channel.Close("主动关闭")
		return
	}

	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),

		"MessageType": messageType,
		"text":        msgText,
	}).Info("OnMessage")

	text := fmt.Sprintf("Receive a message: %s", string(data))
	if err := channel.Push([]byte(text)); err != nil {
		logrus.WithError(err).Error("Fail to push when on message.")
		return
	}
}

func (l *demoListener) BeforeClosedByBackend(channel pushKit.Channel, closeInfo string) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),

		"closeInfo": closeInfo,
	}).Info("BeforeClosedByBackend")
}

func (l *demoListener) OnClose(channel pushKit.Channel, closeInfo string, bsid, user, group string) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),

		"closeInfo": closeInfo,
	}).Info("OnClose")
}
