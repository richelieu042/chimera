package pushKit

import (
	"time"

	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
)

var _ Channel = &BaseChannel{}

// BaseChannel
/*
!!!:
(1) 此类型实现了部分 Channel 接口，
(2) 此类型的子类应当实现 Channel 接口（主要是: Push()、Close()、bind、unbind），不能覆盖父类的方法.
	(由于unbind可能涉及Close()，因此只能在子类中实现)
*/
type BaseChannel struct {
	gmutex.RWMutex

	CloseCh chan string

	// ClientIP 可能是error string（获取失败的情况下）.
	ClientIP string
	// Type Channel 的类型.
	Type string

	Id    string
	Bsid  string
	User  string
	Group string
	Data  interface{}

	Closed    bool
	Listeners Listeners

	PongInterval time.Duration
	Interval     *timeKit.Interval
}

// IsClosed （读锁）
func (channel *BaseChannel) IsClosed() (rst bool) {
	/* 读锁 */
	channel.RLockFunc(func() {
		rst = channel.Closed
	})
	return
}

// SetClosed （写锁）
/*
PS: 返回值如果为true，应当调用 Listeners.OnClose().

@return true: 	设置成功
		false:	设置失败（因为已经被设置关闭）
*/
func (channel *BaseChannel) SetClosed() (flag bool) {
	if channel.Closed {
		return
	}

	/* 写锁 */
	channel.LockFunc(func() {
		if channel.Closed {
			return
		}

		channel.Closed = true
		flag = true
	})
	return
}

// GetCloseCh
/*
	PS: 只有 调用SetClosed() && 返回值为true 的情况下，才能向通道中写数据.
*/
func (channel *BaseChannel) GetCloseCh() chan string {
	return channel.CloseCh
}

func (channel *BaseChannel) GetClientIP() string {
	return channel.ClientIP
}

func (channel *BaseChannel) GetType() string {
	return channel.Type
}

func (channel *BaseChannel) GetId() (id string) {
	// 无需加解锁，因为: id是不变的
	return channel.Id
}

func (channel *BaseChannel) GetBsid() (bsid string) {
	/* 读锁 */
	channel.RLockFunc(func() {
		bsid = channel.Bsid
	})
	return
}

func (channel *BaseChannel) SetBsid(bsid string) {
	/* 写锁 */
	channel.LockFunc(func() {
		channel.Bsid = bsid
	})
}

func (channel *BaseChannel) ClearBsid() {
	channel.SetBsid("")
}

func (channel *BaseChannel) GetUser() (user string) {
	/* 读锁 */
	channel.RLockFunc(func() {
		user = channel.User
	})
	return
}

func (channel *BaseChannel) SetUser(user string) {
	/* 写锁 */
	channel.LockFunc(func() {
		channel.User = user
	})
}

func (channel *BaseChannel) ClearUser() {
	channel.SetUser("")
}

func (channel *BaseChannel) GetGroup() (group string) {
	/* 读锁 */
	channel.RLockFunc(func() {
		group = channel.Group
	})
	return
}

func (channel *BaseChannel) SetGroup(group string) {
	/* 写锁 */
	channel.LockFunc(func() {
		channel.Group = group
	})
}

func (channel *BaseChannel) ClearGroup() {
	channel.SetGroup("")
}

func (channel *BaseChannel) GetData() (data interface{}) {
	/* 读锁 */
	channel.RLockFunc(func() {
		data = channel.Data
	})
	return
}

func (channel *BaseChannel) SetData(data interface{}) {
	/* 写锁 */
	channel.LockFunc(func() {
		channel.Data = data
	})
}

func (channel *BaseChannel) ClearData() {
	channel.SetData(nil)
}

// Equals 根据 id 判断是否相等.
func (channel *BaseChannel) Equals(c Channel) bool {
	if c == nil {
		return false
	}

	return channel.GetId() == c.GetId()
}

func (channel *BaseChannel) Dispose() {
	// 停止主动pong
	channel.Interval.Stop()
	channel.Interval = nil
}
