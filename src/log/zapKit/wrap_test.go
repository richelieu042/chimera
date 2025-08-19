package zapKit

import (
	"fmt"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestWrapLogger(t *testing.T) {
	f, err := os.Create("_a.log")
	if err != nil {
		panic(err)
	}

	encoder := NewEncoder()
	// 确保多个goroutine在写入日志时不会发生竞态条件
	ws := NewLockedWriteSyncer(f)
	core := NewCore(encoder, ws, zapcore.DebugLevel)
	l := zap.New(core)
	wl := WrapLogger(l, f)

	wl.Debug("debug")
	wl.Info("info")
	wl.Warn("warn")
	wl.Error("error")

	fmt.Println("close result:", wl.Close())

	fmt.Println("=== sleep starts ===")
	time.Sleep(time.Second * 3)
	fmt.Println("=== sleep ends ===")
}
