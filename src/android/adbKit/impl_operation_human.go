package adbKit

import "github.com/richelieu042/chimera/v3/src/randomKit"

func (impl *clientImpl) TapAsHumanBeings(x, y int, axisOffset int) error {
	x = x + randomKit.Int(-axisOffset, axisOffset+1)
	y = y + randomKit.Int(-axisOffset, axisOffset+1)

	return impl.Tap(x, y)
}

// LongPressAsHumanBeings
/*
	@param duration 	持续时间（单位：ms）
	@param timeOffset 	持续时间的偏移量（单位：ms）
*/
func (impl *clientImpl) LongPressAsHumanBeings(x, y int, duration int, axisOffset, timeOffset int) error {
	x = x + randomKit.Int(-axisOffset, axisOffset+1)
	y = y + randomKit.Int(-axisOffset, axisOffset+1)

	if duration > 0 {
		duration = duration + randomKit.Int(-timeOffset, timeOffset+1)
	}

	return impl.LongPress(x, y, duration)
}

// SwipeAsHumanBeings 像人一样滑动（每次的位置和时间都不一样）.
/*
	@param duration 	持续时间（单位：ms）
	@param timeOffset 	持续时间的偏移量（单位：ms）
*/
func (impl *clientImpl) SwipeAsHumanBeings(x1, y1, x2, y2 int, duration int, axisOffset, timeOffset int) error {
	x1 = x1 + randomKit.Int(-axisOffset, axisOffset+1)
	y1 = y1 + randomKit.Int(-axisOffset, axisOffset+1)
	x2 = x2 + randomKit.Int(-axisOffset, axisOffset+1)
	y2 = y2 + randomKit.Int(-axisOffset, axisOffset+1)

	if duration > 0 {
		duration = duration + randomKit.Int(-timeOffset, timeOffset+1)
	}

	return impl.Swipe(x1, y1, x2, y2, duration)
}
