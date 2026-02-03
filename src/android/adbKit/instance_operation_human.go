package adbKit

import "github.com/richelieu-yang/chimera/v3/src/randomKit"

func (ins *Instance) TapLikeAHumanBeing(x, y int, axisOffset int) error {
	x = x + randomKit.Int(-axisOffset, axisOffset)
	y = y + randomKit.Int(-axisOffset, axisOffset)

	return ins.Tap(x, y)
}

// LongPressLikeAHumanBeing
/*
	@param duration 	持续时间（单位：ms）
	@param timeOffset 	持续时间的偏移量（单位：ms）
*/
func (ins *Instance) LongPressLikeAHumanBeing(x, y int, duration int, axisOffset, timeOffset int) error {
	x = x + randomKit.Int(-axisOffset, axisOffset)
	y = y + randomKit.Int(-axisOffset, axisOffset)

	if duration > 0 {
		duration = duration + randomKit.Int(-timeOffset, timeOffset)
	}

	return ins.LongPress(x, y, duration)
}

// SwipeLikeAHumanBeing 像人一样滑动（每次的位置和时间都不一样）.
/*
	@param duration 	持续时间（单位：ms）
	@param timeOffset 	持续时间的偏移量（单位：ms）
*/
func (ins *Instance) SwipeLikeAHumanBeing(x1, y1, x2, y2 int, duration int, axisOffset, timeOffset int) error {
	x1 = x1 + randomKit.Int(-axisOffset, axisOffset)
	y1 = y1 + randomKit.Int(-axisOffset, axisOffset)
	x2 = x2 + randomKit.Int(-axisOffset, axisOffset)
	y2 = y2 + randomKit.Int(-axisOffset, axisOffset)

	if duration > 0 {
		duration = duration + randomKit.Int(-timeOffset, timeOffset)
	}

	return ins.Swipe(x1, y1, x2, y2, duration)
}
