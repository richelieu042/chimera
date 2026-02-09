package rocketmq5Kit

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/protocol/v2"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/idKit"
	"github.com/richelieu042/chimera/v3/src/log/logrusKit"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	verifySendTimeout = time.Millisecond * 300

	// verifyTimeLimit verify操作的最长时间，超过就失败
	verifyTimeLimit = time.Second * 3
)

type VerifyConfig struct {
	// Topic 用于验证的topic（理论上，此topic仅用于测试，不能同时用于业务，因为测试发的消息无意义）.
	/*
		PS: 为空则不进行验证.
	*/
	Topic string

	// LogPath 用于测试的日志文件路径（如果为空，则输出到控制台）.
	LogPath string
}

// verify 测试RocketMQ5服务是否正常工作.
/*
@param endpoint 用于测试的RocketMQ5服务的endpoint
@param topic 	用于测试的topic（理论上，此topic仅用于测试，不能同时用于业务，因为测试发的消息无意义）
@return 如果为nil，说明 RocketMQ5服务 正常启动
*/
func verify(config *VerifyConfig) error {
	if config == nil || strKit.IsEmpty(config.Topic) {
		// 不进行验证
		return nil
	}

	/* logger */
	var output io.Writer
	if strKit.IsEmpty(config.LogPath) {
		output = os.Stderr
	} else {
		if err := fileKit.AssertExistAndIsFile(config.LogPath); err != nil {
			return err
		}
		var err error
		output, err = fileKit.CreateInAppendMode(config.LogPath)
		if err != nil {
			return err
		}
	}
	logger := logrusKit.NewLogger(logrusKit.WithOutput(output),
		logrusKit.WithLevel(logrus.DebugLevel),
		logrusKit.WithDisableQuote(true),
		logrusKit.WithMsgPrefix("[RocketMQ5 VERIFY] "),
	)

	topic := config.Topic
	tag := idKit.NewXid()
	consumerGroup := tag
	logger.Infof("topic: [%s]", topic)
	logger.Infof("tag: [%s]", tag)
	logger.Infof("consumer group: [%s]", consumerGroup)

	/* texts */
	timeStr := timeKit.FormatCurrent(timeKit.FormatEntire)
	texts := []string{
		fmt.Sprintf("%s_%s", timeStr, "$0"),
		fmt.Sprintf("%s_%s", timeStr, "$1"),
		fmt.Sprintf("%s_%s", timeStr, "$2"),
		fmt.Sprintf("%s_%s", timeStr, "$3"),
		fmt.Sprintf("%s_%s", timeStr, "$4"),
		fmt.Sprintf("%s_%s", timeStr, "$5"),
	}
	json, err := jsonKit.MarshalIndentToString(texts, "", "    ")
	if err != nil {
		return err
	}
	logger.Infof("texts:\n%s\n.", json)

	/* (1) producer */
	producer, err := NewProducer()
	if err != nil {
		return err
	}
	defer producer.GracefulStop()

	/* (2) consumer */
	consumer, err := NewSimpleConsumer(consumerGroup, map[string]*rmq_client.FilterExpression{
		//topic: rmq_client.SUB_ALL,
		topic: rmq_client.NewFilterExpression(tag),
	})
	if err != nil {
		return err
	}
	defer consumer.GracefulStop()

	ctx, cancel := context.WithTimeout(context.TODO(), verifyTimeLimit)
	defer cancel()
	var producerCh = make(chan error, 1)
	var consumerCh = make(chan error, 1)

	/* (3) producer goroutine */
	go func() {
		defer func() {
			logger.Info("Producer goroutine ends.")
		}()

		for _, text := range texts {
			msg := &rmq_client.Message{
				Topic: topic,
				Body:  []byte(text),
				Tag:   &tag,
			}
			ctx, _ := context.WithTimeout(context.TODO(), verifySendTimeout)
			_, err := producer.Send(ctx, msg)
			if err != nil {
				err = errorKit.Wrapf(err, "Producer fails to send message(%s).", text)
				producerCh <- err
				return
			}
			logger.WithFields(logrus.Fields{
				"text": text,
			}).Info("Producer managers to send a message.")
		}
		logger.Info("Producer managers to send all messages.")
	}()

	/* (4) consumer goroutine */
	textsCopy := sliceKit.Copy(texts)
	go func(texts []string) {
		defer func() {
			logger.Info("Consumer goroutine ends.")
		}()

		for {
			select {
			case <-ctx.Done():
				consumerCh <- errorKit.Newf("Consumer fails to receive all messages within timeout(%s).", verifyTimeLimit)
				return
			case <-time.After(time.Millisecond * 100):
				// do nothing
			}
			//time.Sleep(time.Millisecond * 100)

			mvs, err := consumer.Receive(context.TODO(), DefaultMaxMessageNum, DefaultInvisibleDuration)
			if err != nil {
				/* gRPC errors */
				if s, ok := status.FromError(err); ok {
					switch s.Code() {
					case codes.Canceled:
						consumerCh <- err
						return
						//case codes.Canceled:
						//	/* 提前结束（被取消） */
						//	break LOOP
						//case codes.DeadlineExceeded:
						//	/* 超时结束 */
						//	consumerErr = errorKit.Newf("consumer fails to receive all messages(count: %d) within timeout(%s), missing(%d)", len(texts), verifyTimeLimit.String(), len(text1))
						//	break LOOP
					}
				}

				if errRpcStatus, ok := rmq_client.AsErrRpcStatus(err); ok {
					/* RocketMQ5 errors */
					switch errRpcStatus.Code {
					case int32(v2.Code_MESSAGE_NOT_FOUND):
						// 没有新消息
					default:
						logger.WithFields(logrus.Fields{
							"code":  errRpcStatus.Code,
							"error": err.Error(),
						}).Warn("Consumer fails to receive.")
					}
				} else {
					/* other errors */
					logger.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Warn("Consumer fails to receive.")
				}
				continue
			}

			// ack message
			for _, mv := range mvs {
				text := string(mv.GetBody())

				err := consumer.Ack(context.TODO(), mv)
				if err != nil {
					logger.WithFields(logrus.Fields{
						"tag":   GetTagString(mv.GetTag()),
						"text":  text,
						"error": err.Error(),
					}).Error("Consumer fails to ack message.")
					continue
				}

				var ok bool
				texts, ok = sliceKit.Remove(texts, text)
				left := len(texts)
				logger.WithFields(logrus.Fields{
					"valid": ok,
					"left":  left,
					"tag":   GetTagString(mv.GetTag()),
					"text":  text,
				}).Info("Consumer managers to receive and ack a message.")
				if left == 0 {
					// 成功收到所有预期消息
					logger.Info("Consumer managers to receive and ack all messages.")
					consumerCh <- nil
					return
				}
			}
		}
	}(textsCopy)

	select {
	case producerErr := <-producerCh:
		return producerErr
	case consumerCh := <-consumerCh:
		if consumerCh != nil {
			return consumerCh
		}
		// 通过验证
		return nil
	case <-ctx.Done():
		return errorKit.Newf("Fail to pass verification within timeout(%s).", verifyTimeLimit)
	}
}
