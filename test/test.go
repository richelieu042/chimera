package main

import (
	"github.com/redis/go-redis/v9"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"go.uber.org/zap"
)

func main() {
	err := redis.Nil
	console.Error("出错了", zap.Error(err))

	console.Errorf("出错了1：%v", a())
	console.Errorf("出错了2：%+v", a())
}

func a() error {
	return errorKit.Wrapf(redis.Nil, "wrap")
}
