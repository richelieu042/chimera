package adbKit

import (
	"context"
	"errors"
	"os/exec"

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
	version, cmd, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "version")
	if err != nil {
		return "", "", errKit.Wrapf(err, "fail to run '%s'", cmd.String())
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
	resp, cmd, err := cmdKit.RunCombinedlyToString(context.TODO(), "pkill", "-f", "HD-Adb")
	if err != nil {
		// pkill 返回 exit status 1 表示没有匹配的进程，不是真正的错误
		if exitErr, ok := errors.AsType[*exec.ExitError](err); ok {
			if exitErr.ExitCode() == 1 {
				err = nil // 没有找到进程，忽略
			}
		}
		if err != nil {
			return errKit.Wrapf(err, "fail to run '%s'", cmd)
		}
	}
	logger.Sugar().Debugf("%s:\n%s", cmd, resp)

	// （2）命令：pkill -f adb
	resp, cmd, err = cmdKit.RunCombinedlyToString(context.TODO(), "pkill", "-f", "adb")
	if err != nil {
		return errKit.Wrapf(err, "fail to run '%s'", cmd)
	}
	logger.Sugar().Debugf("%s:\n%s", cmd, resp)

	// （3）命令：adb kill-server
	resp, cmd, err = cmdKit.RunCombinedlyToString(context.TODO(), "adb", "kill-server")
	if err != nil {
		return errKit.Wrapf(err, "fail to run '%s'", cmd)
	}
	logger.Sugar().Debugf("%s:\n%s", cmd, resp)

	// （4）命令：adb start-server
	resp, cmd, err = cmdKit.RunCombinedlyToString(context.TODO(), "adb", "start-server")
	if err != nil {
		return errKit.Wrapf(err, "fail to run '%s'", cmd)
	}
	logger.Sugar().Debugf("%s:\n%s", cmd, resp)

	return nil
}
