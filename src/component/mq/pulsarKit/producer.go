package pulsarKit

import (
	"context"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
)

type (
	Producer struct {
		pulsar.Client
		pulsar.Producer
	}
)

func (p *Producer) SendWithTimeout(pMsg *pulsar.ProducerMessage, timeout time.Duration) (pulsar.MessageID, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	return p.Send(ctx, pMsg)
}

func (p *Producer) Close() {
	if p == nil {
		return
	}

	if p.Producer != nil {
		p.Producer.Close()
	}
	if p.Client != nil {
		p.Client.Close()
	}
}

// NewProducerOriginally
/*
PS: 目标Pulsar服务未启动的情况下，如果ctx不加以限制，要过约 1min 才会返回error（期间客户端日志有connection refused输出）.

@param ctx				建议设置超时时间，否则可能卡死（addresses无效的情况下）
@param options 			必须的属性: Topic
						建议的属性: SendTimeout
@param clientLogPath 	客户端的日志输出（为空则输出到控制台; 不会rotate）
*/
func NewProducerOriginally(ctx context.Context, addresses []string, options pulsar.ProducerOptions, clientLogPath string) (rst *Producer, err error) {
	rst = &Producer{
		Client:   nil,
		Producer: nil,
	}
	defer func() {
		if err != nil {
			err = errorKit.Wrapf(err, "fail to new producer")
		}
	}()

	// case 写入nil: 新建Producer成功
	errCh := make(chan error, 1)

	go func() {
		var err error

		rst.Client, err = NewClient(addresses, clientLogPath)
		if err != nil {
			errCh <- err
			return
		}

		rst.Producer, err = rst.Client.CreateProducer(options)
		if err != nil {
			err = errorKit.Wrapf(err, "client fails to create producer")
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
