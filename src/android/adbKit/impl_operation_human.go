package adbKit

import (
	"context"

	"github.com/richelieu042/chimera/v3/src/randomKit"
)

func (impl *clientImpl) TapAsHumanBeings(ctx context.Context, x, y int, axisOffset int) error {
	x = x + randomKit.Int(-axisOffset, axisOffset+1)
	y = y + randomKit.Int(-axisOffset, axisOffset+1)

	return impl.Tap(ctx, x, y)
}

// LongPressAsHumanBeings
/*
	@param duration 	持续时间（单位：ms）
	@param timeOffset 	持续时间的偏移量（单位：ms）
*/
func (impl *clientImpl) LongPressAsHumanBeings(ctx context.Context, x, y int, duration int, axisOffset, timeOffset int) error {
	x = x + randomKit.Int(-axisOffset, axisOffset+1)
	y = y + randomKit.Int(-axisOffset, axisOffset+1)

	if duration > 0 {
		duration = duration + randomKit.Int(-timeOffset, timeOffset+1)
	}

	return impl.LongPress(ctx, x, y, duration)
}

// SwipeAsHumanBeings 像人一样滑动（每次的位置和时间都不一样）.
/*
	@param duration 	持续时间（单位：ms）
	@param timeOffset 	持续时间的偏移量（单位：ms）
*/
func (impl *clientImpl) SwipeAsHumanBeings(ctx context.Context, x1, y1, x2, y2 int, duration int, axisOffset, timeOffset int) error {
	x1 = x1 + randomKit.Int(-axisOffset, axisOffset+1)
	y1 = y1 + randomKit.Int(-axisOffset, axisOffset+1)
	x2 = x2 + randomKit.Int(-axisOffset, axisOffset+1)
	y2 = y2 + randomKit.Int(-axisOffset, axisOffset+1)

	if duration > 0 {
		duration = duration + randomKit.Int(-timeOffset, timeOffset+1)
	}

	return impl.Swipe(ctx, x1, y1, x2, y2, duration)
}
