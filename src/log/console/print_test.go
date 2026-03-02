package console

import (
	"testing"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func TestPrint(t *testing.T) {
	Debug("DEBUG", zap.String("name", "ZhangSan"))
	Info("INFO")
	Warn("WARN")
	Error("ERROR", zap.Error(redis.Nil))
}
