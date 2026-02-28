package cpuKit

import (
	"bytes"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/shirou/gopsutil/v4/cpu"
)

// GetCpuId 获取cpu的id
/*
PS: cpu.Info() 目前仅支持Windows环境，不支持：Mac环境（M1）、Linux环境.
*/
func GetCpuId() (string, error) {
	stats, err := cpu.Info()
	if err != nil {
		return "", err
	}
	if len(stats) == 0 {
		return "", errorKit.Newf("length of stats is 0")
	}

	buffer := bytes.Buffer{}
	for _, stat := range stats {
		if strKit.IsNotEmpty(stat.PhysicalID) {
			if buffer.Len() > 0 {
				// 多个cpu id间的分隔符
				buffer.WriteString("-")
			}
			buffer.WriteString(stat.PhysicalID)
		}
	}
	if buffer.Len() == 0 {
		return "", errorKit.Newf("length of buffer is 0")
	}
	return buffer.String(), nil
}

//// 通过"wmic命令"来获取.
///*
//Go 获取电脑 CPUID: https://blog.csdn.net/qq_23179075/article/details/83651373
//
//PS:
//微软从 Windows 11 中删除 wmic.exe(https://baijiahao.baidu.com/s?id=1724793484531282024)，
//调用此方法会报error（exec: "wmic": executable file not found in %PATH%）.
//
//golang利用gco获取windows系统cpu信息:
//*/
//func getCpuId0() (string, error) {
//	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
//	out, err := cmd.CombinedOutput()
//	if err == nil {
//		str := string(out)
//		str = strKit.RemoveSpace(str)
//		// e.g.	str: "ProcessorIdBFEBFBFF000906ED"
//		return str[11:], nil
//	}
//	return "", err
//}

//// 通过"C语言代码"来获取.
///*
//golang利用gco获取windows系统cpu信息: https://blog.csdn.net/Man_ge/article/details/111937880
//
//PS: 适用于通过wmic获取cpu id失败的情况.
//*/
//func getCpuId1() (string, error) {
//	var cc *C.char = C.WindowsGetCpuId()
//
//	defer C.free(unsafe.Pointer(cc))
//	str := C.GoString(cc)
//
//	if strKit.IsEmpty(str) {
//		return "", errorKit.Newf("cpu id from C code is empty.")
//	}
//
//	// 通过"-"拆分；反向遍历拼接
//	buffer := bytes.Buffer{}
//	s := strKit.Split(str, "-")
//	for i := len(s) - 1; i >= 0; i-- {
//		buffer.WriteString(s[i])
//	}
//	return buffer.String(), nil
//}
