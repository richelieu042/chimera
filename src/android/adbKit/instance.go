package adbKit

import (
	"context"
	"sync"

	"github.com/richelieu042/chimera/v3/src/command/cmdKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewInstance(address string, cleanFlag, verbose bool) (*Instance, error) {
	enc := zapKit.NewEncoder(zapKit.WithEncoderMessagePrefix("[ABD] "))
	var level zapcore.Level
	if verbose {
		level = zap.InfoLevel
	} else {
		level = zap.ErrorLevel
	}
	core := zapKit.NewCore(enc, nil, level)
	logger := zapKit.NewLogger(core)

	ins := &Instance{
		address:   address,
		cleanFlag: cleanFlag,
		logger:    logger.Sugar(),
	}

	if err := ins.initialize(); err != nil {
		return nil, err
	}
	return ins, nil
}

type Instance struct {
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

func (ins *Instance) initialize() error {
	/* (1) check */
	adbPath, adbVersion, err := Check()
	if err != nil {
		return err
	}
	ins.logger.Info("Check adb environment successfully.", zap.String("path", adbPath), zap.String("version", adbVersion))

	/* (2) clean */
	if ins.cleanFlag {
		if err := Clean(); err != nil {
			return err
		}
		ins.logger.Info("Clean adb environment successfully.")
	}

	// 命令：adb connect {ins.address}
	connectResp, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "connect", ins.address)
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb connect %s'", ins.address)
	}
	if strKit.Index(strKit.ToLower(connectResp), "failed to") != -1 {
		return errorKit.Newf("fail to connect to [%s], response: [%s]", ins.address, connectResp)
	}
	ins.logger.Infof("Connect to [%s] successfully.", ins.address)

	// 命令：adb devices
	devices, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "devices")
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb devices'")
	}
	ins.logger.Infof("adb devices:\n%s", devices)

	return nil
}
