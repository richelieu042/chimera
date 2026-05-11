package adbKit

import (
	"context"
	"regexp"

	"github.com/richelieu042/chimera/v3/src/command/cmdKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/core/intKit"
)

func (impl *clientImpl) GetAddress() string {
	return impl.address
}

// GetPhysicalSize 获取：分辨率（宽高、尺寸）.
func (impl *clientImpl) GetPhysicalSize(ctx context.Context) (width int, height int, err error) {
	/*
		执行命令：adb -s 127.0.0.1:5555 shell wm size
	*/
	str, cmd, err := cmdKit.RunCombinedlyToString(ctx, "adb", "-s", impl.address, "shell", "wm", "size")
	if err != nil {
		return 0, 0, errKit.Wrapf(err, "fail to run '%s'", cmd.String())
	}

	/*
		e.g. "Physical size: 1920x1080"
			matches[0] = "1920x1080" — 完整匹配，即整个正则表达式匹配到的完整字符串
			matches[1] = "1920" — 第一个捕获组，对应第一个括号 (\d+) 匹配的内容
			matches[2] = "1080" — 第二个捕获组，对应第二个括号 (\d+) 匹配的内容
	*/
	re := regexp.MustCompile(`(\d+)x(\d+)$`)
	matches := re.FindStringSubmatch(str)
	if matches == nil {
		return 0, 0, errKit.Newf("fail to get physical size from [%s]", str)
	}
	width, err = intKit.StringToInt(matches[1])
	if err != nil {
		return 0, 0, errKit.Wrapf(err, "fail to get width from [%s]", matches[1])
	}
	height, err = intKit.StringToInt(matches[2])
	if err != nil {
		return 0, 0, errKit.Wrapf(err, "fail to get height from [%s]", matches[2])
	}
	return
}
