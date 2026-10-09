package timeKit

import (
	"context"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestSetInterval(t *testing.T) {
	ctx := context.TODO()
	i, err := SetInterval(ctx, func(t time.Time) {
		logrus.Info(t)
	}, time.Millisecond*300)
	if err != nil {
		panic(err)
	}

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 3)
	logrus.Info("sleep ends")

	ClearInterval(i)
	ClearInterval(i)
	ClearInterval(i)

	logrus.Info("sleep1 starts")
	time.Sleep(time.Second * 3)
	logrus.Info("sleep1 ends")
}

func TestSetInterval1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*1500)
	defer cancel()

	i, err := SetInterval(ctx, func(t time.Time) {
		logrus.Info(t)
	}, time.Second)
	if err != nil {
		panic(err)
	}

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 3)
	logrus.Info("sleep ends")

	ClearInterval(i)
	ClearInterval(i)

	logrus.Info("sleep1 starts")
	time.Sleep(time.Second * 3)
	logrus.Info("sleep1 ends")
}
