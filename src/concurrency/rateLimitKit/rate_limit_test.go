package rateLimitKit

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

func TestNewUberLimiter(t *testing.T) {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	l := log.New(os.Stdout, "", flag)

	limiter := NewUberLimiter(2) // 间隔 500 ms

	limiter.Take() // 第1次 Take：limiter 初始化时允许立即取，几乎不等待，立即返回
	l.Println("1")
	limiter.Take() // 第2次 Take：距上次 Take 不足 500ms，等待补足，约等 ~500ms 后返回
	l.Println("2")
	limiter.Take() // 第3次 Take：距上次 Take 不足 500ms，等待补足，约等 ~500ms 后返回
	l.Println("3")
}

// 默认（有 slack）的例子
func TestNewUberLimiter1(t *testing.T) {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	l := log.New(os.Stdout, "", flag)

	limiter := NewUberLimiter(2) // 2 RPS，间隔 500ms，默认有 slack

	limiter.Take() // t=0，立即返回
	l.Println("0")

	time.Sleep(time.Second * 1) // 这里 sleep 了 1 秒（积累了 slack）

	limiter.Take()
	l.Println("1") // 立即返回（消耗积累的 slack）
	limiter.Take()
	l.Println("2") // 立即返回（还有剩余 slack）
	limiter.Take()
	l.Println("3") // 等待 ~500ms（slack 耗尽）
}

// WithoutSlack 的例子
func TestNewUberLimiter2(t *testing.T) {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	l := log.New(os.Stdout, "", flag)

	limiter := NewUberLimiter(2, ratelimit.WithoutSlack) // 2 RPS，间隔 500ms，没有 slack

	limiter.Take() // t=0，立即返回
	l.Println("0")

	time.Sleep(time.Second * 1) // 空闲再久也不积累

	limiter.Take()
	l.Println("1") // 立即返回（距上次超过 500ms，无需等待）
	limiter.Take()
	l.Println("2") // 等待 ~500ms（严格间隔，slack 不起作用）
	limiter.Take()
	l.Println("3") // // 等待 ~500ms
}

func TestNewLimiter(t *testing.T) {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	l := log.New(os.Stdout, "", flag)

	limiter := NewLimiter(rate.Every(500*time.Millisecond), 1)

	limiter.Wait(context.Background())
	l.Println("1")
	limiter.Wait(context.Background())
	l.Println("2")
	limiter.Wait(context.Background())
	l.Println("3")
}
