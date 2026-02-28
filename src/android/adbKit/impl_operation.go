package adbKit

import (
	"context"
	"os"

	"github.com/richelieu042/chimera/v3/src/command/cmdKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/intKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/randomKit"
)

// Screenshot 截图
/*
	@param targetPath 截图保存的路径（PNG格式）
*/
func (impl *clientImpl) Screenshot(targetPath string) error {
	impl.Lock()
	defer impl.Unlock()

	/*
		执行命令：adb -s 127.0.0.1:5555 exec-out screencap -p
		-p: 参数表示以 PNG 格式输出截图数据
	*/
	data, err := cmdKit.RunCombinedly(context.TODO(), "adb", "-s", impl.address, "exec-out", "screencap", "-p")
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb -s %s exec-out screencap -p'", impl.address)
	}

	// 尝试创建父目录
	if err := fileKit.MkParentDirs(targetPath); err != nil {
		return errorKit.Wrapf(err, "fail to make parent dirs of [%s]", targetPath)
	}

	err = os.WriteFile(targetPath, data, 0644)
	if err != nil {
		return errorKit.Wrapf(err, "fail to write file [%s]", targetPath)
	}

	return nil
}

// Tap 点击.
/*
	命令：adb -s 127.0.0.1:5555 shell input tap <x> <y>
*/
func (impl *clientImpl) Tap(x, y int) error {
	resp, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "-s", impl.address, "shell", "input", "tap", intKit.IntToString(x), intKit.IntToString(y))
	if err != nil {
		return err
	}
	if strKit.IsNotEmpty(resp) {
		return errorKit.Newf("fail to tap, response: %s", resp)
	}
	return nil
}

// LongPress 长按.
/*
	命令：adb -s 127.0.0.1:5555 shell input swipe 500 1000 500 1000 2000

	@param duration: 持续时间（单位：ms）
*/
func (impl *clientImpl) LongPress(x, y int, duration int) error {
	if duration <= 0 {
		duration = randomKit.Int(300, 401) // 默认: 300-400ms
	}

	resp, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "-s", impl.address, "shell", "input", "swipe", intKit.IntToString(x), intKit.IntToString(y), intKit.IntToString(x), intKit.IntToString(y),
		intKit.IntToString(duration),
	)
	if err != nil {
		return err
	}
	if strKit.IsNotEmpty(resp) {
		return errorKit.Newf("fail to tap, response: %s", resp)
	}
	return nil
}

// Swipe 滑动.
/*
	命令：adb -s 127.0.0.1:5555 shell input swipe <x1> <y1> <x2> <y2> <duration>

	@param x1, y1: 起始坐标
	@param x2, y2: 结束坐标
	@param duration: 持续时间（单位：ms）

	e.g. 	向上滑动（上滑刷新/滚动）
		adb -s 127.0.0.1:5555 shell input swipe 500 1500 500 500 300
	e.g.1 	向下滑动（下拉通知栏）
		adb -s 127.0.0.1:5555 shell input swipe 500 100 500 1000 300
	e.g.2 	向左滑动
		adb -s 127.0.0.1:5555 shell input swipe 900 500 100 500 300
	e.g.3 	向右滑动
		adb -s 127.0.0.1:5555 shell input swipe 100 500 900 500 300
*/
func (impl *clientImpl) Swipe(x1, y1, x2, y2 int, duration int) error {
	if duration <= 0 {
		duration = randomKit.Int(300, 401) // 默认: 300-400ms
	}

	resp, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "-s", impl.address, "shell", "input", "swipe",
		intKit.IntToString(x1), intKit.IntToString(y1),
		intKit.IntToString(x2), intKit.IntToString(y2),
		intKit.IntToString(duration),
	)
	if err != nil {
		return err
	}
	if strKit.IsNotEmpty(resp) {
		return errorKit.Newf("fail to tap, response: %s", resp)
	}
	return nil
}
