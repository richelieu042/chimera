package adbKit

import (
	"context"
	"regexp"

	"github.com/richelieu042/chimera/v3/src/command/cmdKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/intKit"
)

// GetPhysicalSize 获取：分辨率（宽高、尺寸）.
func (ins *Instance) GetPhysicalSize() (width int, height int, err error) {
	/*
		执行命令：adb -s 127.0.0.1:5555 shell wm size
	*/
	str, err := cmdKit.RunCombinedlyToString(context.TODO(), true, "adb", "-s", ins.address, "shell", "wm", "size")
	if err != nil {
		return 0, 0, errorKit.Wrapf(err, "fail to run 'adb -s %s shell wm size'", ins.address)
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
		return 0, 0, errorKit.Newf("fail to get physical size from [%s]", str)
	}
	width, err = intKit.StringToInt(matches[1])
	if err != nil {
		return 0, 0, errorKit.Wrapf(err, "fail to get width from [%s]", matches[1])
	}
	height, err = intKit.StringToInt(matches[2])
	if err != nil {
		return 0, 0, errorKit.Wrapf(err, "fail to get height from [%s]", matches[2])
	}
	return
}
