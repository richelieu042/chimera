package pushKit

import (
	"errors"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

var (
	NotSetupError = errKit.Newf("haven’t been set up correctly")

	ChannelClosedError = errKit.Newf("channel has already been closed")

	NoSuitableChannelError = errKit.Newf("no suitable channel")
)

// IsNoSuitableChannelError 推送返回的error，是否是因为不存在对应的channel？
func IsNoSuitableChannelError(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, NoSuitableChannelError)
}
