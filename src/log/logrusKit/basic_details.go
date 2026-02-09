package logrusKit

import (
	"github.com/richelieu042/chimera/v3/src/log/commonLogKit"
	"github.com/sirupsen/logrus"
)

// PrintBasicDetails 输出服务器的基本信息（以便于甩锅）
func PrintBasicDetails(logger *logrus.Logger) {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	commonLogKit.PrintBasicDetails(logger)
}
