package otelKit

import (
	"github.com/richelieu042/chimera/v3/src/atomic/atomicKit"
	"go.uber.org/atomic"
)

var setupFlag *atomic.Bool = atomicKit.NewBool(false)

func check() error {
	if !setupFlag.Load() {
		return NotSetupError
	}
	return nil
}
