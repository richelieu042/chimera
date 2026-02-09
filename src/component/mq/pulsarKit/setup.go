package pulsarKit

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

var config *Config

// MustSetUp
/*
PS: Pulsar服务中途挂掉的话，恢复后，Consumer实例、Producer实例还能正常工作.

@param verifyConfig		可以为nil（将不验证，但不推荐这么干）
*/
func MustSetUp(pulsarConfig *Config, verifyConfig *VerifyConfig) {
	if err := SetUp(pulsarConfig, verifyConfig); err != nil {
		console.Fatalf("Fail to set up, error: %s", err.Error())
	}
}

func SetUp(pulsarConfig *Config, verifyConfig *VerifyConfig) (err error) {
	defer func() {
		if err != nil {
			config = nil
		}
	}()

	if err = validateKit.Struct(pulsarConfig); err != nil {
		err = errorKit.Wrapf(err, "Fail to verify")
		return
	}
	config = pulsarConfig

	// verify
	if err = verify(verifyConfig); err != nil {
		err = errorKit.Wrapf(err, "Fail to verify")
		return
	}

	return
}

// NewProducer
/*
前提: 成功调用 SetUp() || MustSetUp().

@param options 必需属性: Topic、SendTimeout
@param logPath 客户端的日志输出（"": 输出到控制台）
*/
func NewProducer(ctx context.Context, options pulsar.ProducerOptions, clientLogPath string) (*Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	return NewProducerOriginally(ctx, config.Addrs, options, clientLogPath)
}

// NewConsumer
/*
前提: 成功调用 SetUp() || MustSetUp().

@param options 必需属性: Topic、SubscriptionName、Type
@param logPath 客户端的日志输出（"": 输出到控制台）
*/
func NewConsumer(ctx context.Context, options pulsar.ConsumerOptions, clientLogPath string) (*Consumer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	return NewConsumerOriginally(ctx, config.Addrs, options, clientLogPath)
}
