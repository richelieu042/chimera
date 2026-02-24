package adbKit

import (
	"go.uber.org/zap"
)

type Client interface {
	GetAddress() string

	GetPhysicalSize() (width int, height int, err error)

	Screenshot(targetPath string) error

	// Tap 点击.
	Tap(x, y int) error
	// LongPress 长按.
	LongPress(x, y int, duration int) error
	// Swipe 滑动.
	Swipe(x1, y1, x2, y2 int, duration int) error

	TapAsHumanBeings(x, y int, axisOffset int) error
	LongPressAsHumanBeings(x, y int, duration int, axisOffset, timeOffset int) error
	SwipeAsHumanBeings(x1, y1, x2, y2 int, duration int, axisOffset, timeOffset int) error
}

// NewClient
/*
@param logger: 可以为nil（默认：丢弃输出）
*/
func NewClient(address string, cleanFlag bool, logger *zap.SugaredLogger) (Client, error) {
	ins := &clientImpl{
		address:   address,
		cleanFlag: cleanFlag,
	}

	if err := ins.initialize(logger); err != nil {
		return nil, err
	}
	return ins, nil
}
