package slbKit

import (
	"os"

	"github.com/richelieu042/chimera/v3/src/atomic/atomicKit"
	"github.com/richelieu042/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLoadBalancer
/*
PS: 返回的*LoadBalancer实例，需要手动调用 Start 以启动.

@param logger 可以为nil，默认输出到控制台
*/
func NewLoadBalancer(logger *zap.Logger) (lb *LoadBalancer) {
	if logger == nil {
		encoder := zapKit.NewEncoder(zapKit.WithEncoderMessagePrefix("[SLB] "))
		ws := os.Stdout
		core := zapKit.NewCore(encoder, ws, zapcore.DebugLevel)

		logger = zapKit.NewLogger(core, zapKit.WithAddStacktrace(zapcore.DPanicLevel))
	}

	lb = &LoadBalancer{
		logger:   logger,
		backends: nil,
		current:  atomicKit.NewInt64(-1),
		status:   StatusInitialized,
	}
	return
}
