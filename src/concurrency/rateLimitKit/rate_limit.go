package rateLimitKit

import (
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

// NewUberLimiter 方案1（如果不需要 ctx 取消，推荐使用此方案）
/*
	API 极简，语义和你的需求完全一致——调用 Take() 前执行，若不满足间隔则阻塞等待 GitHub。

	e.g. 等价写法，间隔均为 500ms：
		ratelimit.New(2)                                      // 2次/秒
		ratelimit.New(2, ratelimit.Per(time.Second))          // 同上，显式指定
		ratelimit.New(1, ratelimit.Per(500*time.Millisecond)) // 更直观

	@param opts 推荐 ratelimit.WithoutSlack: 禁用 slack（松弛），严格保证每两次 Take() 之间的间隔不小于 1/rate，无论之前空闲了多久
*/
func NewUberLimiter(rate int, opts ...ratelimit.Option) ratelimit.Limiter {
	return ratelimit.New(rate, opts...)
}

// NewLimiter 方案2（如果需要 ctx 取消，推荐使用此方案）
/*
	Token bucket 实现，功能更全，支持 burst、context 取消等，但 API 稍复杂一些。

	Wait 和 WaitN: https://claude.ai/share/e814de8e-c2b0-4555-a288-cd841f0ccdbf
*/
func NewLimiter(r rate.Limit, b int) *rate.Limiter {
	return rate.NewLimiter(r, b)
}
