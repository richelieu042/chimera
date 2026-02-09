package rocketmq5Kit

import (
	"context"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
)

// NewProducer
/*
PS:
(1) In most case, you don't need to create many producers, singletion pattern is more recommended.
(2) 需要先 set up!!!
(3) 第一个返回值，非nil的情况下，不再需要时请调用 GracefulStop().

@param clientLogPath 客户端日志（blank则输出到控制台）
*/
func NewProducer() (rmq_client.Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	endpoint := sliceKit.Join(config.Endpoints, ";")
	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint:    endpoint,
		Credentials: config.Credentials,
	})
	if err != nil {
		return nil, errorKit.Wrapf(err, "Fail to new producer")
	}
	if err := producer.Start(); err != nil {
		return nil, errorKit.Wrapf(err, "Fail to start producer")
	}
	return producer, nil
}

// SendMessage Deprecated: 仅供参考如何生产消息
func SendMessage(producer rmq_client.Producer, ctx context.Context, topic string, tag *string, body []byte, messageGroup string, deliveryTimestamp time.Time, keys ...string) ([]*rmq_client.SendReceipt, error) {
	msg := &rmq_client.Message{
		Topic: topic,
		Tag:   tag,
		Body:  body,
	}
	msg.SetKeys(keys...)
	msg.SetMessageGroup(messageGroup)
	msg.SetDelayTimestamp(deliveryTimestamp)
	return producer.Send(ctx, msg)
}
