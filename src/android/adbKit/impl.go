package adbKit

import (
	"context"
	"sync"

	"github.com/richelieu042/chimera/v3/src/command/cmdKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"go.uber.org/zap"
)

type clientImpl struct {
	Client

	/*
		截图的锁，防止并发冲突.
		PS: 截图频率不要太高，建议至少要间隔500~1000ms。
	*/
	sync.Mutex

	// address 安卓设备的地址
	/*
		（1）基本语法：adb connect <设备IP地址>:<端口号>
		（2）默认端口号是 5555，如果使用默认端口，可以省略端口号
			e.g. adb connect 192.168.1.100 <=> adb connect 192.168.1.100:5555
		（3）mac上，BlueStacks Air默认是 "127.0.0.1:5555"
	*/
	address string

	cleanFlag bool

	logger *zap.SugaredLogger
}

func (impl *clientImpl) initialize() error {
	/* (1) check */
	adbPath, adbVersion, err := Check()
	if err != nil {
		return err
	}
	impl.logger.Info("Check adb environment successfully.", zap.String("path", adbPath), zap.String("version", adbVersion))

	/* (2) clean */
	if impl.cleanFlag {
		if err := Clean(); err != nil {
			return err
		}
		impl.logger.Info("Clean adb environment successfully.")
	}

	// 命令：adb connect {impl.address}
	connectResp, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "connect", impl.address)
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb connect %s'", impl.address)
	}
	if strKit.Index(strKit.ToLower(connectResp), "failed to") != -1 {
		return errorKit.Newf("fail to connect to [%s], response: [%s]", impl.address, connectResp)
	}
	impl.logger.Infof("Connect to [%s] successfully.", impl.address)

	// 命令：adb devices
	devices, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "devices")
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb devices'")
	}
	impl.logger.Infof("adb devices:\n%s", devices)

	return nil
}
