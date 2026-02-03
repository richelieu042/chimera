package adbKit

import (
	"context"
	"regexp"
	"sync"

	"github.com/richelieu-yang/chimera/v3/src/command/cmdKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/intKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
)

func NewInstance(address string, cleanFlag bool) *Instance {
	return &Instance{
		address:   address,
		cleanFlag: cleanFlag,
	}
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
}

func (ins *Instance) CheckEnv() error {
	path, err := cmdKit.LookPath("adb")
	if err != nil {
		return errorKit.Wrapf(err, "fail to look path of adb")
	}
	console.Infof("adb path: [%s]", path)

	// adb 版本号
	{
		str, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "version")
		if err != nil {
			return errorKit.Wrapf(err, "fail to run 'adb version'")
		}
		console.Infof("adb version:\n%s", str)
	}

	return nil
}

func (ins *Instance) Initialize() error {
	if ins.cleanFlag {
		// 命令：pkill -f HD-Instance
		// Richelieu: 此处返回的 err 不用管
		_, _ = cmdKit.RunCombinedlyToString(context.TODO(), "pkill", "-f", "HD-Instance")

		// 命令：pkill -f adb
		// Richelieu: 此处返回的 err 不用管
		_, _ = cmdKit.RunCombinedlyToString(context.TODO(), "pkill", "-f", "adb")

		// 命令：adb kill-server
		_, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "kill-server")
		if err != nil {
			return errorKit.Wrapf(err, "fail to run 'adb kill-server'")
		}

		// 命令：adb start-server
		_, err = cmdKit.RunCombinedlyToString(context.TODO(), "adb", "start-server")
		if err != nil {
			return errorKit.Wrapf(err, "fail to run 'adb start-server'")
		}
	}

	// 命令：adb connect {ins.address}
	_, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "connect", ins.address)
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb connect %s'", ins.address)
	}
	console.Infof("Connect to [%s] successfully.", ins.address)

	// adb devices
	devices, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "devices")
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb devices'")
	}
	console.Infof("adb devices:\n%s", devices)

	return nil
}

// GetPhysicalSize 获取：分辨率（宽高、尺寸）.
func (ins *Instance) GetPhysicalSize() (width int, height int, err error) {
	/*
		执行命令：adb -s 127.0.0.1:5555 shell wm size
	*/
	str, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "-s", ins.address, "shell", "wm", "size")
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
