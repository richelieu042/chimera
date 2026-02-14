package adbKit

type Client interface {
	GetPhysicalSize() (width int, height int, err error)

	Screenshot(targetPath string) error

	Tap(x, y int) error
	LongPress(x, y int, duration int) error
	Swipe(x1, y1, x2, y2 int, duration int) error

	TapLikeAHumanBeing(x, y int, axisOffset int) error
	LongPressLikeAHumanBeing(x, y int, duration int, axisOffset, timeOffset int) error
	SwipeLikeAHumanBeing(x1, y1, x2, y2 int, duration int, axisOffset, timeOffset int) error
}
