package main

import (
	"github.com/cockroachdb/errors"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"go.uber.org/zap"
)

func main() {
	console.Error("hello, world", zap.Error(redis.Nil))

	err := errorKit.Newf("123")
	console.Error("hello, world", zap.Error(err))

	{
		err := errors.New("456")
		console.Error("hello, world", zap.Error(err))
		console.Errorf("hello, world: %+v", err)
	}

	errors.NewWithDepth()
}
