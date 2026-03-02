package errKit

import (
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"go.uber.org/zap"

	goerrors "errors"
)

func TestNew(t *testing.T) {
	err := New("ccc")

	console.Error("hello world", zap.Error(err))
	console.Info("---")
	console.Errorf("hello world: %v", err)
	console.Info("---")
	console.Errorf("hello world: %+v", err)
}

func TestWrap(t *testing.T) {
	err := Wrap(redis.Nil, "ccc")

	console.Error("hello world", zap.Error(err))
	console.Info("---")
	console.Errorf("hello world: %v", err)
	console.Info("---")
	console.Errorf("hello world: %+v", err)
}

func TestAs(t *testing.T) {
	err := Wrap(redis.Nil, "ccc")
	fmt.Println(goerrors.As(err, redis.Nil))
	//fmt.Println(As(redis.Nil, err))
}
