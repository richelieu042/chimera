//go:build !windows

package osKit

import (
	"context"
	"strconv"

	"github.com/richelieu042/chimera/v3/src/command/cmdKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/intKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// GetUlimitInfo 获取: 目前资源限制的信息.
/*
命令: sh -c "ulimit -a"
*/
func GetUlimitInfo() (string, error) {
	data, _, err := cmdKit.RunCombinedly(context.TODO(), "sh", "-c", "ulimit -a")
	if err != nil {
		return "", err
	}
	str := strKit.TrimSpace(string(data))

	return str, nil
}

// GetMaxOpenFiles 获取: 同一时间最多可开启的文件数.
/*
PS:
(1) 当前仅支持Mac、Linux环境.
(2) 为何使用 sh -c "ulimit -n" 而非 ulimit -n? https://www.thinbug.com/q/17483723
*/
func GetMaxOpenFiles() (int, error) {
	data, _, err := cmdKit.RunCombinedly(context.TODO(), "sh", "-c", "ulimit -n")
	if err != nil {
		return 0, err
	}
	// e.g. "122880\n" => "122880"
	str := strKit.TrimSpace(string(data))

	i, err := intKit.StringToInt(str)
	if err != nil {
		return 0, errorKit.Newf("result(%s) isn't a number", str)
	}
	return i, nil

	//cmd := exec.Command("sh", "-c", "ulimit -n")
	////cmd := exec.Command("ulimit", "-n")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	return 0, err
	//}
	//// strKit.TrimSpace()是为了：去掉最后面的"\n"
	//str := strKit.TrimSpace(out.String())
	//value, err := strconv.Atoi(str)
	//if err != nil {
	//	return 0, errorKit.Newf("result(%s) of command(%s) isn't a number", str, cmd.String())
	//}
	//return value, nil
}

// GetMaxProcessThreadCountByUser 获取: 单个用户可以创建的进程数上限（线程也算）
/*
PS:
(1) ulimit -u命令: 限制单个用户可以创建的进程数.
(2) ulimit -u命令也可以用来限制单个用户可以创建的线程数，因为: 在Linux中，线程本质上只是具有共享地址空间的进程。
*/
func GetMaxProcessThreadCountByUser() (int, error) {
	data, _, err := cmdKit.RunCombinedly(context.TODO(), "sh", "-c", "ulimit -u")
	if err != nil {
		return 0, err
	}
	// e.g."5333\n" => "5333"
	str := strKit.TrimSpace(string(data))

	i, err := intKit.StringToInt(str)
	if err != nil {
		return 0, errorKit.Newf("result(%s) isn't a number", str)
	}
	return i, nil
}

// GetCoreFileSize 获取: core文件的最大值，单位为区块.
func GetCoreFileSize() (string, error) {
	data, _, err := cmdKit.RunCombinedly(context.TODO(), "sh", "-c", "ulimit -c")
	if err != nil {
		return "", err
	}
	// e.g."5333\n" => "5333"
	str := strKit.TrimSpace(string(data))

	if str == "unlimited" {
		return str, nil
	}
	i, err := intKit.StringToInt(str)
	if err != nil {
		return "", errorKit.Newf("result(%s) isn't a number", str)
	}
	return strconv.Itoa(i), nil
}
