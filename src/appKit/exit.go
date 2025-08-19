package appKit

import (
	"time"

	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
)

var (
	// Exit 退出程序（应用）.
	Exit func(codes ...int) = zapKit.Exit

	RegisterExitHandler func(handlers ...func()) = zapKit.RegisterExitHandler

	RegisterParallelExitHandler func(handlers ...func()) = zapKit.RegisterParallelExitHandler

	// SetExitTimeout 执行所有exit handler的超时时间.
	SetExitTimeout func(d time.Duration) = zapKit.SetExitTimeout

	RunExitHandlers func() = zapKit.RunExitHandlers
)
