package signalKit

import (
	"os"
	"os/signal"

	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
)

//var printOnce sync.Once
//var exitOnce sync.Once
//
//// MonitorExitSignals 异步地监听退出信号（拦截关闭信号）.
///*
//可以参考 go-zero 中的 "core/proc/signals.go".
//
//PS:
//(1) 无法拦截部分信号（e.g. syscall.SIGSTOP、syscall.SIGKILL）;
//(2) 可以通过 logrus.RegisterExitHandler() 在程序退出前"毁尸灭迹"（在里面你甚至可以 time.Sleep）;
//(3) 此函数对 主动调用os.Exit() 无效;
//(4) 信号处理函数中不要使用 fmt.Println 等函数，因为它们不是线程安全的，会导致程序崩溃;
//(5) 虽然可以多次调用本函数，但不推荐这么干，1次就够了.
//*/
//func MonitorExitSignals(callback func(sig os.Signal)) {
//	ch := make(chan os.Signal, 1)
//	signal.Notify(ch, ExitSignals...)
//
//	go func() {
//		sig := <-ch
//		printOnce.Do(func() {
//			logrus.Warnf("Receive an exit signal(%s).", sig.String())
//		})
//
//		callback(sig)
//		exitOnce.Do(func() {
//			time.Sleep(time.Second * 3)
//
//			logrus.Fatalf("Process exits with signal(%s).", sig.String())
//		})
//	}()
//}

// MonitorExitSignals 同步地监听.
/*
@param callbacks 可以不传

PS:
(1) 会 阻塞 调用此函数的goroutine;
(2) 理论上，应该 由main goroutine调用此函数 && 此函数只能被调用1次.
*/
func MonitorExitSignals() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, ExitSignals...)

	sig := <-ch
	console.Warnf("Receive an exit signal(%s).", sig.String())

	zapKit.Exit(1)
}

//// runCallback 防止执行callback时发生 panic（参考了logrus中的runHandler）.
//func runCallback(sig os.Signal, callback func(sig os.Signal)) {
//	defer func() {
//		if err := recover(); err != nil {
//			logrus.WithField("err", err).Error("Recover from execute callback.")
//		}
//	}()
//
//	callback(sig)
//}
