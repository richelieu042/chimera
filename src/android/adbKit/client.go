package adbKit

import (
	"context"

	"go.uber.org/zap"
)

type Client interface {
	GetAddress() string

	GetPhysicalSize(ctx context.Context) (width int, height int, err error)

	Screenshot(ctx context.Context, targetPath string) error

	// Tap 点击.
	Tap(ctx context.Context, x, y int) error
	// LongPress 长按.
	LongPress(ctx context.Context, x, y int, duration int) error
	// Swipe 滑动.
	Swipe(ctx context.Context, x1, y1, x2, y2 int, duration int) error

	TapAsHumanBeings(ctx context.Context, x, y int, axisOffset int) error
	LongPressAsHumanBeings(ctx context.Context, x, y int, duration int, axisOffset, timeOffset int) error
	SwipeAsHumanBeings(ctx context.Context, x1, y1, x2, y2 int, duration int, axisOffset, timeOffset int) error
}

// NewClient
/*
@param logger: 可以为nil（默认：丢弃输出）
*/
func NewClient(ctx context.Context, address string, cleanFlag bool, logger *zap.Logger) (Client, error) {
	if logger == nil {
		logger = zap.NewNop() // 不输出
	}

	ins := &clientImpl{
		address:   address,
		cleanFlag: cleanFlag,
	}
	if err := ins.initialize(ctx, logger); err != nil {
		return nil, err
	}
	return ins, nil
}
