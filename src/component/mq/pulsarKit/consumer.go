package pulsarKit

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
)

type (
	Consumer struct {
		pulsar.Client
		pulsar.Consumer
	}
)

func (c *Consumer) Close() {
	if c == nil {
		return
	}

	if c.Consumer != nil {
		c.Consumer.Close()
	}
	if c.Client != nil {
		c.Client.Close()
	}
}

// NewConsumerOriginally
/*
PS: 目标Pulsar服务未启动的情况下，如果ctx不加以限制，要过约 1min 才会返回error（期间客户端日志有connection refused输出）.

@param ctx				建议设置超时时间，否则可能卡死（addresses无效的情况下）
@param options			必须的属性: Topic、SubscriptionName、Type
@param clientLogPath 	客户端的日志输出（为空则输出到控制台; 不会rotate）
*/
func NewConsumerOriginally(ctx context.Context, addresses []string, options pulsar.ConsumerOptions, clientLogPath string) (rst *Consumer, err error) {
	rst = &Consumer{
		Client:   nil,
		Consumer: nil,
	}
	defer func() {
		if err != nil {
			err = errorKit.Wrapf(err, "fail to new consumer")
		}
	}()

	// case 写入nil: 新建Consumer成功
	errCh := make(chan error, 1)

	go func() {
		var err error

		rst.Client, err = NewClient(addresses, clientLogPath)
		if err != nil {
			errCh <- err
			return
		}

		rst.Consumer, err = rst.Client.Subscribe(options)
		if err != nil {
			err = errorKit.Wrapf(err, "client fails to subscribe with topic(%s), subscriptionName(%s) and type(%s)",
				options.Topic, options.SubscriptionName, options.Type)
			errCh <- err
			return
		}

		errCh <- nil
	}()

	select {
	case <-ctx.Done():
		rst.Close()
		return nil, ctx.Err()
	case err := <-errCh:
		if err != nil {
			rst.Close()
			return nil, err
		}
		return rst, nil
	}
}
