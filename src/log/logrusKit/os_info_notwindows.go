//go:build !windows

package logrusKit

import (
	"github.com/richelieu042/chimera/v3/src/core/osKit"
	"github.com/sirupsen/logrus"
)

// printUniqueOsInfo 输出特殊的信息（主要依赖于不同的OS）
func printUniqueOsInfo() {
	if str, err := osKit.GetUlimitInfo(); err != nil {
		logrus.WithError(err).Warn("[CHIMERA, OS] fail to get ulimit information")
	} else {
		DisableQuoteTemporarily(nil, func(logger *logrus.Logger) {
			logger.Infof("[CHIMERA, OS] ulimit information:\n%s\n", str)
		})
	}

	if i, err := osKit.GetThreadsMax(); err != nil {
		logrus.WithError(err).Warn("[CHIMERA, OS] fail to get kernel.threads-max")
	} else {
		logrus.Infof("[CHIMERA, OS] kernel.threads-max: [%d].", i)
	}
	if i, err := osKit.GetPidMax(); err != nil {
		logrus.WithError(err).Warn("[CHIMERA, OS] fail to get kernel.pid_max")
	} else {
		logrus.Infof("[CHIMERA, OS] kernel.pid_max: [%d].", i)
	}
	if i, err := osKit.GetMaxMapCount(); err != nil {
		logrus.WithError(err).Warn("[CHIMERA, OS] fail to get vm.max_map_count")
	} else {
		logrus.Infof("[CHIMERA, OS] vm.max_map_count: [%d].", i)
	}
}
