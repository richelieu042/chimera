package timeKit

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gmutex"
)

// Interval
/*
参考:
Golang正确停止Ticker https://blog.csdn.net/weixin_40098405/article/details/111517279

!!!:
(1) 使用 time.Ticker 时要注意: Stop会停止Ticker，停止后，Ticker不会再被发送，但是Stop不会关闭通道，防止读取通道发生错误;
(2) 可以重复调用 Ticker.Stop().
*/
type Interval struct {
	gmutex.RWMutex

	stopped bool

	ticker *time.Ticker

	// closeCh 关闭通道.
	closeCh chan struct{}
}

// Stop
/*
PS:
(1) 可以多次调用，不会panic，但这样没意义（只有第一次调用才有意义）;
(2) 如果有任务正在执行，会等它先执行完.
*/
func (i *Interval) Stop() {
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

// SetInterval 效果类似于JavaScript中的 window.setInterval().
/*
@param ctx		控制 *Interval实例 的生命周期
@param task		(1) 不能为nil
				(2) 传参t为执行任务时的 time.Time
@param duration 必须>0
*/
func SetInterval(ctx context.Context, task func(t time.Time), duration time.Duration) *Interval {
	i := &Interval{
		stopped: false,
		ticker:  time.NewTicker(duration),
		closeCh: make(chan struct{}),
	}

	go func(i *Interval) {
		//// test
		//defer func() {
		//	logrus.Info("[TEST] goroutine ends...")
		//}()

		defer i.ticker.Stop()

		for {
			select {
			case t := <-i.ticker.C:
				if i.IsStopped() {
					// 中断的是最内层的select语句
					break
				}
				task(t)
			case <-i.closeCh:
				return
			case <-ctx.Done():
				i.Stop()
				return
			}
		}
	}(i)
	return i
}

// ClearInterval 效果类似于JavaScript中的同名函数.
/*
@param i 可以为nil
*/
func ClearInterval(i *Interval) {
	i.Stop()
}
