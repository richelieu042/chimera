package pushKit

import (
	"github.com/panjf2000/ants/v2"
	"github.com/richelieu042/chimera/v3/src/atomic/atomicKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/validateKit"
	"github.com/sirupsen/logrus"
)

var (
	setupFlag = atomicKit.NewBool(false)

	// pushPool 并发执行推送任务
	pushPool *ants.Pool
)

func MustSetUp(antPool *ants.Pool, logger *logrus.Logger) {
	if err := Setup(antPool, logger); err != nil {
		console.Fatalf("Fail to set up, error: %s", err.Error())
	}
}

// Setup
/*
@param antPool	用来并发执行推送任务
				(1) 不能为nil
				(2) 需要自行决定: cap大小、是否自定义输出...
@param logger 	可以为nil
@param options
*/
func Setup(antPool *ants.Pool, logger *logrus.Logger) (err error) {
	defer func() {
		setupFlag.Store(err == nil)
	}()

	/* antPool */
	if err := interfaceKit.AssertNotNil(antPool, "antPool"); err != nil {
		return err
	}
	if antPool.IsClosed() {
		return errorKit.Newf("antPool has already been closed")
	}
	capacity := antPool.Cap()
	if capacity > 0 {
		tag := "gte=100"
		if err := validateKit.Var(capacity, tag); err != nil {
			return errorKit.Wrapf(err, "capacity(%d) of antPool is invalid(tag: %s) when it's greater than zero", capacity, tag)
		}
	}
	pushPool = antPool

	/* logger */
	if logger != nil {
		if err := SetDefaultLogger(logger); err != nil {
			return err
		}
	}

	return nil
}

func CheckSetup() error {
	if !setupFlag.Load() {
		return NotSetupError
	}
	return nil
}
