package pulsarKit

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu042/chimera/v3/src/core/conditionKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/idKit"
	"github.com/richelieu042/chimera/v3/src/log/logrusKit"
	"github.com/richelieu042/chimera/v3/src/randomKit"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
	"github.com/sirupsen/logrus"
)

const (
	// connectTimeout 创建Consumer（或Producer）的超时时间
	connectTimeout = time.Second * 10

	// receiveTimeout 接受消息的超时时间
	receiveTimeout = time.Second * 10

	// sendTimeout 单次发送消息的超时时间
	sendTimeout = time.Second
)

// verify 简单地验证 Pulsar服务 是否启动成功
/*
PS:
可能失败的原因：
（1）pulsar的进程在，但启动报错（存储空间爆了）
*/
func verify(config *VerifyConfig) (err error) {
	if config == nil || strKit.IsBlank(config.Topic) {
		// 不验证
		return nil
	}

	ulid := idKit.NewUlid()

	// 对应客户端日志s生成在 临时目录 下
	tmpDirPath, err := pathKit.GetExclusiveTempDir()
	if err != nil {
		return err
	}
	timeStr := timeKit.FormatCurrent(timeKit.FormatFileName)
	consumerLogPath := pathKit.Join(tmpDirPath, fmt.Sprintf("pulsar_verify_consumer_%s_%s.log", timeStr, ulid))
	producerLogPath := pathKit.Join(tmpDirPath, fmt.Sprintf("pulsar_verify_producer_%s_%s.log", timeStr, ulid))

	// 是否打印日志到控制台？
	level := conditionKit.TernaryOperator(config.Print, logrus.DebugLevel, logrus.PanicLevel)
	logger := logrusKit.NewLogger(logrusKit.WithLevel(level), logrusKit.WithMsgPrefix("[PULSAR, VERIFY] "))
	logger.Infof("consumerLogPath: [%s].", consumerLogPath)
	logger.Infof("producerLogPath: [%s].", producerLogPath)

	defer func() {
		if err == nil {
			// 验证成功的情况下，删掉客户端日志文件
			if err := fileKit.RemoveAll(consumerLogPath); err != nil {
				logger.WithError(err).Error("Fail to delete consumerLogPath.")
			} else {
				logger.Info("Delete consumerLogPath.")
			}
			if err := fileKit.RemoveAll(producerLogPath); err != nil {
				logger.WithError(err).Error("Fail to delete producerLogPath.")
			} else {
				logger.Info("Delete producerLogPath.")
			}
		}
	}()

	err = _verify(logger, config.Topic, consumerLogPath, producerLogPath, ulid)
	return
}

func _verify(logger *logrus.Logger, topic, consumerLogPath, producerLogPath, ulid string) error {
	ctx0, cancel := context.WithTimeout(context.TODO(), connectTimeout)
	defer cancel()
	consumer, err := NewConsumer(ctx0, pulsar.ConsumerOptions{
		Topic:                       topic,
		SubscriptionName:            fmt.Sprintf("verify_%s_%d", ulid, randomKit.Int(0, 10000000)),
		Type:                        pulsar.Exclusive,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionLatest,
	}, consumerLogPath)
	if err != nil {
		return err
	}
	defer consumer.Close()

	ctx1, cancel := context.WithTimeout(context.TODO(), connectTimeout)
	defer cancel()
	producer, err := NewProducer(ctx1, pulsar.ProducerOptions{
		Topic:       topic,
		SendTimeout: sendTimeout,
	}, producerLogPath)
	if err != nil {
		return err
	}
	defer producer.Close()

	timeStr := timeKit.FormatCurrent(timeKit.FormatEntire)
	texts := []string{
		fmt.Sprintf("%s&&%s&&%s", timeStr, ulid, "$0"),
		fmt.Sprintf("%s&&%s&&%s", timeStr, ulid, "$1"),
		fmt.Sprintf("%s&&%s&&%s", timeStr, ulid, "$2"),
		fmt.Sprintf("%s&&%s&&%s", timeStr, ulid, "$3"),
		fmt.Sprintf("%s&&%s&&%s", timeStr, ulid, "$4"),
		fmt.Sprintf("%s&&%s&&%s", timeStr, ulid, "$5"),
	}
	var ch = make(chan struct{}, 1)
	var consumerErrCh = make(chan error, 1)
	var producerErrCh = make(chan error, 1)

	/* consumer */
	consumerCtx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	go func() {
		defer func() {
			logger.Info("(CONSUMER) Goroutine ends.")
		}()

		s := sliceKit.Copy(texts)
		for {
			msg, err := consumer.Receive(consumerCtx)
			if err != nil {
				logger.WithError(err).Info("(CONSUMER) Fail to receive.")
				consumerErrCh <- err
				break
			}
			if err := consumer.Ack(msg); err != nil {
				logger.WithError(err).Info("(CONSUMER) Fail to ack.")
				consumerErrCh <- err
				break
			}

			var ok bool
			text := string(msg.Payload())
			s, ok = sliceKit.Remove(s, text)
			left := len(s)
			logger.WithFields(logrus.Fields{
				"left":  left,
				"valid": ok,
				"text":  text,
			}).Info("(CONSUMER) Receive a message.")

			if ok && left == 0 {
				logger.Info("(CONSUMER) Receive all messages!")
				ch <- struct{}{}
				break
			}
		}
	}()

	/* producer */
	go func() {
		defer func() {
			logger.Info("(PRODUCER) Goroutine ends.")
		}()

		for _, text := range texts {
			pMsg := &pulsar.ProducerMessage{
				Payload: []byte(text),
			}
			err := func() error {
				ctx, cancel := context.WithTimeout(context.TODO(), sendTimeout)
				defer cancel()
				_, err := producer.Send(ctx, pMsg)
				return err
			}()
			if err != nil {
				logger.WithFields(logrus.Fields{
					"text":  text,
					"error": err.Error(),
				}).Error("(PRODUCER) Fail to send.")
				producerErrCh <- err
				break
			}
			logger.WithFields(logrus.Fields{
				"text": text,
			}).Info("(PRODUCER) Manager to send.")
		}
	}()

	select {
	case <-ch:
		return nil
	case err := <-producerErrCh:
		return err
	case err := <-consumerErrCh:
		return err
	case <-time.After(receiveTimeout):
		return errorKit.Newf("Fail to get all messages within timeout(%s)", receiveTimeout)
	}
}
