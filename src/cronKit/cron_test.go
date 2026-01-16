package cronKit

import (
	"log"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestNewCronWithTask(t *testing.T) {
	c, _, err := NewCronWithTask("@every 3s", func() {
		logrus.Info("do")
	})
	if err != nil {
		panic(err)
	}
	c.Start()

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")

	// 可以多次调用 Cron.Stop()，虽然只有第一次有意义，但至少不会panic
	logrus.Info("0")
	select {
	case <-c.Stop().Done():
	}
	logrus.Info("1")
	select {
	case <-c.Stop().Done():
	}
	logrus.Info("2")

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 4)
	logrus.Info("sleep ends")
}

// 通过 Cron.Stop 停止任务.
func TestStopCron(t *testing.T) {
	c, _, err := NewCronWithTask("@every 3s", func() {
		log.Println("do")
	})
	if err != nil {
		panic(err)
	}
	c.Start()

	time.Sleep(time.Second * 10)
	StopCron(c)
	log.Println("---")

	// 再等会，看任务到底停了没
	time.Sleep(time.Second * 10)
	log.Println("===")
}

// 通过 Cron.Remove 停止任务.
func TestRemoveCron(t *testing.T) {
	c, entryId, err := NewCronWithTask("@every 3s", func() {
		log.Println("do")
	})
	if err != nil {
		panic(err)
	}
	c.Start()

	time.Sleep(time.Second * 10)
	c.Remove(entryId)
	log.Println("---")

	// 再等会，看任务到底停了没
	time.Sleep(time.Second * 10)
	log.Println("===")
}
