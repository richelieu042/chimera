package main

import (
	"time"

	"github.com/richelieu-yang/chimera/v3/src/android/adbKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/randomKit"
	"go.uber.org/zap"
)

func main() {
	ins := adbKit.NewInstance("127.0.0.1:5555", true, true)
	if err := ins.Initialize(); err != nil {
		console.Fatalf("fail to initialize: %+v", err)
	}

	count := 0
	for {
		count++

		minute := randomKit.RandFloat(0.5, 1.5, 2) * 1000 * 60
		d := time.Millisecond * time.Duration(minute)
		console.Info("Sleep starts.", zap.Int("count", count), zap.String("duration", d.String()))
		time.Sleep(d)
		console.Info("Sleep ends.", zap.Int("count", count), zap.String("duration", d.String()))

		if err := ins.Swipe(500, 1500, 500, 500, 300); err != nil {
			console.Error("Fail to swipe.", zap.Error(err))
			continue
		}
		console.Info("Manager to swipe.")
	}
}
