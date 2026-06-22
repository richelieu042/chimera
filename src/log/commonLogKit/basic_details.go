package commonLogKit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/cpuKit"
	"github.com/richelieu042/chimera/v3/src/core/memoryKit"
	"github.com/richelieu042/chimera/v3/src/core/osKit"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
	"github.com/richelieu042/chimera/v3/src/core/runtimeKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/core/userKit"
	"github.com/richelieu042/chimera/v3/src/dataSizeKit"
	"github.com/richelieu042/chimera/v3/src/diskKit"
	"github.com/richelieu042/chimera/v3/src/ip/ipKit"
	"github.com/richelieu042/chimera/v3/src/processKit"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
)

func PrintBasicDetails(logger Logger) {
	if logger == nil {
		return
	}

	logger.Info(strings.Repeat("=", 42))
	logger.Infof("\n%s", consts.Banner)

	// 作用: 通知外部（可以是多个）协程a已执行完毕
	chA := make(chan struct{})
	// 作用: 通知外部（可以是多个）协程b已执行完毕
	chB := make(chan struct{})

	/* 协程a */
	go func() {
		logger.Infof("[CHIMERA, PROCESS] pid: [%d]", processKit.PID)

		/* golang */
		logger.Infof("[CHIMERA, GO] version: [%s]", runtimeKit.GetGoVersion())
		logger.Infof("[CHIMERA, GO] GOROOT: [%s]", runtimeKit.GetGoRoot())

		/* os */
		logger.Infof("[CHIMERA, OS] os: [%s]", osKit.OS)
		logger.Infof("[CHIMERA, OS] arch: [%s]", osKit.ARCH)
		logger.Infof("[CHIMERA, OS] bits: [%d]", osKit.GetOsBits())
		printUlimitInfo(logger)
		printLimitsForCurrentPid(logger)
		printOsInfo(logger)
		printCgroupInfo(logger)

		/* user */
		logger.Infof("[CHIMERA, USER] uid: [%s]", userKit.GetUid())
		logger.Infof("[CHIMERA, USER] gid: [%s]", userKit.GetGid())
		logger.Infof("[CHIMERA, USER] name: [%s]", userKit.GetName())
		logger.Infof("[CHIMERA, USER] user name: [%s]", userKit.GetUserName())
		logger.Infof("[CHIMERA, USER] home dir: [%s]", userKit.GetUserHomeDir())

		/* path */
		logger.Infof("[CHIMERA, PATH] working directory: [%s]", pathKit.GetWorkingDir())
		logger.Infof("[CHIMERA, PATH] os temporary directory: [%s]", pathKit.GetOsTempDir())
		logger.Infof("[CHIMERA, PATH] self dir: [%s]", pathKit.SelfDir())
		logger.Infof("[CHIMERA, PATH] main pkg path: [%s]", pathKit.MainPkgPath())

		/* json */
		logger.Infof("[CHIMERA, JSON] library: [%s]", jsonKit.GetLibrary())

		/* ip */
		logger.Infof("[CHIMERA, IP] internal ip: [%s]", ipKit.GetInternalIp())
		ips := ipKit.GetIps()
		logger.Infof("[CHIMERA, IP] ips: [%s]", sliceKit.Join(ips, ", "))

		/* host */
		if hostInfo, err := runtimeKit.GetHostInfo(); err != nil {
			logger.Warnf("[CHIMERA, HOST] fail to get host info, error: %s", err.Error())
		} else {
			logger.Infof("[CHIMERA, HOST] host name: [%s]", hostInfo.Hostname)
		}

		/* CPU */
		printCpuDetails(logger)

		/* memory */
		printMemoryDetails(logger)

		/* disk */
		printDiskDetails(logger)

		// 关闭信道，通知外部（可以是多个）: 协程a 已执行完毕
		close(chA)
	}()

	/* 协程b */
	go func() {
		/* time */
		printTimeDetails(logger, chA)

		// 关闭信道，通知外部（可以是多个）: 协程b 已执行完毕
		close(chB)
	}()

	// 等待 协程a 执行完毕
	<-chA

	select {
	case <-chB:
		// 协程b 执行完毕
	case <-time.After(time.Millisecond * 200):
		// 等了200ms（但 协程b 还在执行，就不管它了）
	}

	logger.Info(strings.Repeat("=", 42))
}

func printTimeDetails(logger Logger, ch chan struct{}) {
	// 最多3s
	reqCtx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	networkTime, source, err := timeKit.GetNetworkTime(reqCtx)
	machineTime := timeKit.GetMachineTime()
	zoneName, zoneOffset := machineTime.Zone()

	// 等待 协程a 执行完毕，防止: 多协程输出导致输出混在一起，很难看
	<-ch

	if err != nil {
		logger.Warnf("[CHIMERA, TIME] Fail to get network time, error: %s", err.Error())
	} else {
		logger.Infof("[CHIMERA, TIME] network time: [%v], source: [%s]", networkTime, source)
	}
	logger.Infof("[CHIMERA, TIME] machine time: [%v], zone: [%s, %d]", machineTime, zoneName, zoneOffset)
}

func printMemoryDetails(logger Logger) {
	stats, err := memoryKit.GetMachineMemoryStat()
	if err != nil {
		logger.Errorf("[CHIMERA, MEMORY] fail to get machine memory stats, error: %s", err.Error())
		return
	}
	str := fmt.Sprintf("total: %s, available: %s, used percent: %.2f%%",
		dataSizeKit.ToReadableIecString(float64(stats.Total)),
		dataSizeKit.ToReadableIecString(float64(stats.Available)),
		stats.UsedPercent,
	)
	logger.Infof("[CHIMERA, MEMORY] machine memory stats: [%s]", str)
}

func printDiskDetails(logger Logger) {
	stats, err := diskKit.GetDiskUsageStats()
	if err != nil {
		logger.Warnf("[CHIMERA, DISK] fail to get disk usage stats, error: %s", err.Error())
	} else {
		str := fmt.Sprintf("path: %s, free: %s, used: %s, total: %s, used percent: %.2f%%",
			stats.Path,
			dataSizeKit.ToReadableIecString(float64(stats.Free)),
			dataSizeKit.ToReadableIecString(float64(stats.Used)),
			dataSizeKit.ToReadableIecString(float64(stats.Total)),
			stats.UsedPercent,
		)
		logger.Infof("[CHIMERA, DISK] disk usage stats: [%s]", str)
	}
}

func printCpuDetails(logger Logger) {
	logger.Infof("[CHIMERA, CPU] in a virtual machine? [%t]", cpuKit.InVirtualMachine())
	logger.Infof("[CHIMERA, CPU] vendor id: [%s]", cpuKit.GetVendorID())
	logger.Infof("[CHIMERA, CPU] vendor string: [%s]", cpuKit.GetVendorString())
	logger.Infof("[CHIMERA, CPU] brand name: [%s]", cpuKit.GetBrandName())
	logger.Infof("[CHIMERA, CPU] CPU number: [%d]", cpuKit.GetCpuNumber())
	logger.Infof("[CHIMERA, CPU] features: [%s]", sliceKit.Join(cpuKit.GetFeatureSet(), ","))
	logger.Infof("[CHIMERA, CPU] frequency: [%d]hz", cpuKit.GetFrequency())

	usage, err := cpuKit.GetUsagePercent()
	if err != nil {
		logger.Warnf("[CHIMERA, CPU] fail to get uasge percent, error: %s", err.Error())
	} else {
		logger.Infof("[CHIMERA, CPU] uasge percent: [%.2f]%%", usage)
	}
}
