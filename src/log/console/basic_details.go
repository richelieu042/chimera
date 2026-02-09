package console

import (
	"github.com/richelieu042/chimera/v3/src/log/commonLogKit"
)

func PrintBasicDetails() {
	logger := S()
	defer logger.Sync()

	commonLogKit.PrintBasicDetails(logger)
}
