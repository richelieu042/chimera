package pushKit

import (
	"errors"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
)

var (
	NotSetupError = errorKit.Newf("haven’t been set up correctly")

	ChannelClosedError = errorKit.Newf("channel has already been closed")

	NoSuitableChannelError = errorKit.Newf("no suitable channel")
)

// IsNoSuitableChannelError 推送返回的error，是否是因为不存在对应的channel？
func IsNoSuitableChannelError(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, NoSuitableChannelError)
}
