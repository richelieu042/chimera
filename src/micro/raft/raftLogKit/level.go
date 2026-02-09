package raftLogKit

import (
	"github.com/hashicorp/go-hclog"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

func StringToLevel(levelStr string) hclog.Level {
	if strKit.IsEmpty(levelStr) {
		return hclog.Debug
	}

	return hclog.LevelFromString(levelStr)
}
