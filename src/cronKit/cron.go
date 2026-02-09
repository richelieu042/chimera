package cronKit

import (
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/robfig/cron/v3"
)

// NewCron
/*
!!!:
(1) 想要通过修改机器时间来验证的话，需要先改时间，再启动cron.
(2) 返回的 *cron.Cron实例 要调用 Run() || Start() 以"启动"
	(a) Run()	会阻塞 调用此方法的goroutine，
	(b) Start()	不会阻塞 调用此方法的goroutine
(3) 希望停止任务，可以调用 cron.Cron 的 Stop() || Remove(id EntryID).

!!!: 在线Cron表达式生成器
	https://cron.qqe2.com/
定时任务-表达式
	https://goframe.org/pages/viewpage.action?pageId=30736411
Go 每日一库之定时任务库：cron
	https://mp.weixin.qq.com/s/swdijAro2k8LuYu7q_La1A
cron表达式，每天凌晨0点执行定时任务
	https://www.cnblogs.com/yddwinter/p/16033633.html

e.g.
"* * * * * ?"			每秒，执行一次
"0 0 * * * ?"			每小时，执行一次
"0/2 * * * * ?"			每2秒，执行一次
"0 0/2 * * * ?"			每2分钟，执行一次
"0 0 10,14,16 * * ?"	每天上午10点、下午2点、4点，执行一次
"0 0/30 9-17 * * ?"		朝九晚五工作时间内每半小时，执行一次
"0 30 21 * * ?"			每天晚上21:30，执行一次
"15,30 * * * * ?"		每分钟的第15s、第30s，执行一次

e.g.1
"0 0 0/3 * * ?"			从 0小时 开始，每 3小时 执行一次
"@every 10s"			从启动（调用Run() || Start()）开始，每 10s	执行一次
"@every 1m"				从启动（调用Run() || Start()）开始，每 1min	执行一次
"@hourly"				从启动（调用Run() || Start()）开始，每 1h 	执行一次
"@every 1h30m"			从启动（调用Run() || Start()）开始，每 1.5h 执行一次
*/
func NewCron() *cron.Cron {
	// cron.WithSeconds(): 带"秒"
	return cron.New(cron.WithSeconds())
}

// NewCronWithTask
/*
@param spec "@every 10s" || "@every 1m"，更多可参考"Golang - 1.docx"
*/
func NewCronWithTask(spec string, task func()) (c *cron.Cron, entryId cron.EntryID, err error) {
	defer func() {
		if err != nil {
			StopCron(c)
			c = nil
			entryId = 0
		}
	}()

	if err = strKit.AssertNotBlank(spec, "spec"); err != nil {
		return
	}
	if err = interfaceKit.AssertNotNil(task, "task"); err != nil {
		return
	}

	c = NewCron()
	entryId, err = c.AddFunc(spec, task)
	if err != nil {
		return
	}
	return
}

func NewCronWithJob(spec string, job cron.Job) (c *cron.Cron, entryId cron.EntryID, err error) {
	defer func() {
		if err != nil {
			StopCron(c)
			c = nil
			entryId = 0
		}
	}()

	if err = strKit.AssertNotBlank(spec, "spec"); err != nil {
		return
	}
	if err = interfaceKit.AssertNotNil(job, "job"); err != nil {
		return
	}

	c = NewCron()
	entryId, err = c.AddJob(spec, job)
	if err != nil {
		return
	}
	return
}

// StopCron
/*
@param c 	(1) 可以为nil;
			(2) 可以是 未启动(Start || Run) 的*cron.Cron实例;
			(3) 可以是 已经停止(Stop) 的*cron.Cron实例.

!!!:
(1) 调用此函数可能会 阻塞 调用的goroutine;
(2) 可以多次调用 Cron.Stop()，虽然只有第一次有意义，但至少不会panic.
*/
func StopCron(c *cron.Cron) {
	if c == nil {
		return
	}

	ctx := c.Stop()
	// 如果存在正在执行的任务，会阻塞直到它完成
	select {
	case <-ctx.Done():
	}
}
