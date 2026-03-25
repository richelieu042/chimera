package adbKit

import (
	"context"

	"github.com/richelieu042/chimera/v3/src/command/cmdKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
)

// Check 检查 adb 环境.
/*
@param logger: 可以为nil（默认输出到控制台）
@return path: adb可执行文件的绝对路径
@return version: adb版本号
*/
func Check() (path string, version string, err error) {
	path, err = cmdKit.LookPath("adb")
	if err != nil {
		return "", "", errKit.Wrapf(err, "fail to look path of adb")
	}

	// adb 版本号
	version, err = cmdKit.RunCombinedlyToString(context.TODO(), "adb", "version")
	if err != nil {
		return "", "", errKit.Wrapf(err, "fail to run 'adb version'")
	}

	return
}

// Clean 清理 adb 环境.
/*
!!!: 调用此函数前，需要先调用 Check.
*/
func Clean(logger *zap.Logger) error {
	if logger == nil {
		logger = zapKit.NewNopLogger() // 不记录
	}

	// （1）命令：pkill -f HD-Adb
	// Richelieu: 此处返回的 err 不用管
	_, err := cmdKit.RunCombinedlyToString(context.TODO(), "pkill", "-f", "HD-Adb")
	if err != nil {
		logger.Sugar().Warnf("fail to run 'pkill -f HD-Adb', error: %+v", err)
	}

	// （2）命令：pkill -f adb
	// Richelieu: 此处返回的 err 不用管
	_, err = cmdKit.RunCombinedlyToString(context.TODO(), "pkill", "-f", "adb")
	if err != nil {
		logger.Sugar().Warnf("fail to run 'pkill -f adb', error: %+v", err)
	}

	// （3）命令：adb kill-server
	_, err = cmdKit.RunCombinedlyToString(context.TODO(), "adb", "kill-server")
	if err != nil {
		return errKit.Wrapf(err, "fail to run 'adb kill-server'")
	}

	// （4）命令：adb start-server
	_, err = cmdKit.RunCombinedlyToString(context.TODO(), "adb", "start-server")
	if err != nil {
		return errKit.Wrapf(err, "fail to run 'adb start-server'")
	}

	return nil
}
