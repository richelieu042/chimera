package main

import (
	"github.com/redis/go-redis/v9"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"go.uber.org/zap"
)

func main() {
	// （1）简单error
	err := redis.Nil
	console.Error("出错了", zap.Error(err))

	console.Info("---------")

	// （2）带堆栈信息的error
	console.Errorf("出错了1：%v", a())
	console.Errorf("出错了2：%+v", a())
}

func a() error {
	return errorKit.Wrapf(redis.Nil, "wrap")
}
