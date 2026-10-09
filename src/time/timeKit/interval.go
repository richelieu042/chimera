package timeKit

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

// Interval 表示由 SetInterval 创建的周期任务。
// 使用 Stop 停止并等待任务退出；任务回调内部需要停止时使用 StopAsync。
// Interval 的零值不可用，创建后不可复制。
type Interval struct {
	gmutex.RWMutex

	stopped bool

	ticker *time.Ticker

	// closeCh 在请求停止时关闭，用于通知工作 goroutine 退出。
	closeCh chan struct{}

	// doneCh 在工作 goroutine 退出后关闭，用于让 Stop 等待任务结束。
	doneCh chan struct{}
}

// Stop 请求停止，并等待正在执行的回调和工作 goroutine 退出。
// 可重复调用；nil 接收者会直接返回。
// 不要在任务回调中调用 Stop，否则会等待当前回调结束而死锁；此时应调用 StopAsync。
func (i *Interval) Stop() {
	if i == nil {
		return
	}

	i.StopAsync()
	<-i.doneCh
}

// StopAsync 请求停止，但不等待正在执行的回调结束。
// 可在任务回调中调用，也可重复调用；nil 接收者会直接返回。
// 返回后，通过停止检查的回调仍可能开始或继续执行；如需等待，应从回调外调用 Stop。
func (i *Interval) StopAsync() {
	if i == nil {
		return
	}

	/* 写锁 */
	i.LockFunc(func() {
		if i.stopped {
			return
		}

		i.stopped = true
		i.ticker.Stop()
		close(i.closeCh)
	})
}

// IsStopped 报告是否已请求停止，不表示正在执行的回调已经结束。
// nil 接收者返回 true。
func (i *Interval) IsStopped() (rst bool) {
	if i == nil {
		return true
	}

	/* 读锁 */
	i.RLockFunc(func() {
		rst = i.stopped
	})
	return
}

// SetInterval 每隔 duration 触发一次 task，并返回用于停止任务的 Interval。
// ctx 和 task 不能为 nil，duration 必须大于零，否则返回 nil 和错误。
// 第一次触发发生在一个周期之后；执行较慢时 ticker 可能丢弃部分 tick。
// task 在同一个工作 goroutine 中串行执行；传入的时间是 ticker 的触发时间，
// 不一定是回调实际开始执行的时间。取消 ctx 会请求停止，但不会中断正在执行的回调。
func SetInterval(ctx context.Context, task func(t time.Time), duration time.Duration) (*Interval, error) {
	if ctx == nil {
		return nil, errKit.New("timeKit.SetInterval: nil context")
	}
	if task == nil {
		return nil, errKit.New("timeKit.SetInterval: nil task")
	}
	if duration <= 0 {
		return nil, errKit.New("timeKit.SetInterval: duration must be greater than zero")
	}

	ctxDone := ctx.Done()
	i := &Interval{
		ticker:  time.NewTicker(duration),
		closeCh: make(chan struct{}),
		doneCh:  make(chan struct{}),
	}

	go func(i *Interval) {
		defer close(i.doneCh)
		defer i.StopAsync()

		for {
			select {
			case t := <-i.ticker.C:
				if ctx.Err() != nil {
					return
				}
				if i.IsStopped() {
					return
				}
				task(t)
			case <-i.closeCh:
				return
			case <-ctxDone:
				return
			}
		}
	}(i)
	return i, nil
}

// ClearInterval 停止任务并等待退出，等同于 i.Stop()；i 可以为 nil。
// 不要在任务回调中调用；回调内部应使用 i.StopAsync()。
func ClearInterval(i *Interval) {
	i.Stop()
}
