package commonLogKit

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/richelieu042/chimera/v3/src/command/cmdKit"
	"github.com/richelieu042/chimera/v3/src/core/osKit"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

func printUlimitInfo(logger Logger) {
	/*
		Richelieu: 针对 Ubuntu，使用 bash 而 非sh，因为默认情况下 /bin/sh -> dash*.
	*/
	//cmd := exec.CommandContext(context.TODO(), "sh", "-c", "ulimit -a")
	cmd := exec.CommandContext(context.TODO(), "bash", "-c", "ulimit -a")
	data, err := cmd.CombinedOutput()
	if err != nil {
		logger.Warnf("[CHIMERA, OS] command(%s) fails, error: %s", cmd.String(), err.Error())
		return
	}
	str := strings.TrimSpace(string(data))
	logger.Infof("[CHIMERA, OS] ulimit imformation(ulimit -a):\n%s", str)
}

// printLimitsForCurrentPid
/*
关于 ulimit 的两个天坑
	https://mp.weixin.qq.com/s/yUwQW3pnBhN50tHwBLsiMw
*/
func printLimitsForCurrentPid(logger Logger) {
	pid := os.Getpid()

	cmd := exec.CommandContext(context.TODO(), "bash", "-c", fmt.Sprintf("cat /proc/%d/limits", pid))
	data, err := cmd.CombinedOutput()
	if err != nil {
		logger.Warnf("[CHIMERA, OS] command(%s) fails, error: %s", cmd.String(), err.Error())
		return
	}
	str := strings.TrimSpace(string(data))
	logger.Infof("[CHIMERA, OS] limit for pid(%d):\n%s", pid, str)
}

func printOsInfo(logger Logger) {
	var kernelParameterKeys = []string{
		"fs.file-max",
		"fs.nr_open",
		"fs.inotify.max_user_watches",
		"fs.aio-max-nr",
		"fs.epoll.max_user_watches",
		"fs.pipe-max-size",

		"kernel.pid_max",
		"kernel.threads-max",
		"kernel.hung_task_timeout_secs",
		"kernel.sched_latency_ns",         // Centos 7支持，Ubuntu 24.04 LTS不支持
		"kernel.sched_migration_cost_ns",  // Centos 7支持，Ubuntu 24.04 LTS不支持
		"kernel.sched_min_granularity_ns", // Centos 7支持，Ubuntu 24.04 LTS不支持
		"kernel.timer_migration",
		"kernel.numa_balancing",

		"net.core.netdev_max_backlog",
		"net.core.somaxconn",
		"net.core.rmem_max",
		"net.core.wmem_max",
		"net.core.rmem_default",
		"net.core.wmem_default",

		"net.ipv4.ip_local_port_range",
		"net.ipv4.ip_forward",
		"net.ipv4.tcp_rmem",
		"net.ipv4.tcp_wmem",
		"net.ipv4.tcp_fin_timeout",
		"net.ipv4.tcp_max_syn_backlog",
		//"net.ipv4.tcp_tw_recycle", // 被 net.ipv4.tcp_tw_reuse 替代
		"net.ipv4.tcp_tw_reuse",
		"net.ipv4.tcp_mtu_probing",
		"net.ipv4.udp_mem",
		"net.ipv4.tcp_max_syn_backlog",
		"net.ipv4.tcp_max_tw_buckets",

		"vm.dirty_background_ratio",
		"vm.dirty_ratio",
		"vm.max_map_count",
		"vm.overcommit_memory",
		"vm.swappiness",
		"vm.vfs_cache_pressure",
	}

	// getKernelParameterValue 获取Linux内核参数的值.
	var getKernelParameterValue = func(key string) (string, error) {
		s := strKit.Split(key, ".")
		path := pathKit.Join("/proc/sys", strKit.Join(s, "/"))

		cmd := exec.CommandContext(context.TODO(), "bash", "-c", fmt.Sprintf("cat %s", path))
		data, err := cmd.CombinedOutput()
		if err != nil {
			return "", err
		}
		str := strings.TrimSpace(string(data))
		return str, nil
	}

	for _, key := range kernelParameterKeys {
		value, err := getKernelParameterValue(key)
		if err != nil {
			logger.Warnf("[CHIMERA, OS] %s: fail to get, error: %s", key, err.Error())
			continue
		}
		logger.Infof("[CHIMERA, OS] %s: %s", key, value)
	}
}

// printCgroupInfo
/*
命令（stat -fc %T /sys/fs/cgroup/）用于判断系统当前使用的是 cgroup v1 还是 cgroup v2:
(1) 如果输出为 tmpfs，则表示系统使用的是 cgroup v1;
(2) 如果输出为 cgroup2fs，则表示系统使用的是 cgroup v2.
*/
func printCgroupInfo(logger Logger) {
	cgroupType, _, err := cmdKit.RunToString(context.TODO(), "bash", "-c", "stat -fc %T /sys/fs/cgroup/")
	if err != nil {
		logger.Warnf("Fail to get cgroup type, error: %s", err.Error())
		return
	}
	logger.Infof("[CHIMERA, OS] cgroup type: %s", cgroupType)

	var softLimitPath, hardLimitPath string
	if cgroupType == "cgroup2fs" {
		/* cgroup v2 */
		cgroupPath, err := getCgroupPath()
		if err != nil {
			logger.Warnf("Fail to get cgroup path, error: %s", err.Error())
			return
		}
		softLimitPath = pathKit.Join(cgroupPath, "memory.min") // v2 中没有直接对应的软限制
		hardLimitPath = pathKit.Join(cgroupPath, "memory.max")
	} else {
		/* cgroup v1 */
		softLimitPath = "/sys/fs/cgroup/memory/memory.soft_limit_in_bytes"
		hardLimitPath = "/sys/fs/cgroup/memory/memory.limit_in_bytes"
	}

	printLine := func(path string) {
		s := strKit.Split(path, osKit.PathSeparator)
		if len(s) <= 3 {
			logger.Warnf("[CHIMERA, OS] invalid path: %s", path)
			return
		}
		key := s[len(s)-1]

		cmd := exec.Command("bash", "-c", fmt.Sprintf("cat %s", path))
		data, err := cmd.CombinedOutput()
		if err != nil {
			logger.Warnf("[CHIMERA, OS] Command(%s) fails, error: %s", cmd.String(), err.Error())
			return
		}
		value := strKit.TrimSpace(string(data))
		logger.Infof("[CHIMERA, OS] %s: %s", key, value)
	}
	printLine(softLimitPath)
	printLine(hardLimitPath)
}

// getCgroupPath 获取当前进程的 cgroup 路径（cgroup v2）
func getCgroupPath() (string, error) {
	file, err := os.Open("/proc/self/cgroup")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// cgroup2 的格式为 0::/path/to/cgroup
		parts := strings.Split(line, ":")
		if len(parts) == 3 && parts[0] == "0" {
			// 获取 cgroup 路径
			return filepath.Join("/sys/fs/cgroup", parts[2]), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("cgroup path not found")
}
