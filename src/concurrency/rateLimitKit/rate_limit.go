package rateLimitKit

import (
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

// NewUberLimiter 方案1（如果不需要 ctx 取消，推荐使用此方案）
/*
	API 极简，语义和你的需求完全一致——调用 Take() 前执行，若不满足间隔则阻塞等待 GitHub。
*/
func NewUberLimiter(rate int, opts ...ratelimit.Option) ratelimit.Limiter {
	return ratelimit.New(rate, opts...)
}

// NewLimiter 方案2（如果需要 ctx 取消，推荐使用此方案）
/*
	Token bucket 实现，功能更全，支持 burst、context 取消等，但 API 稍复杂一些。
*/
func NewLimiter(r rate.Limit, b int) *rate.Limiter {
	return rate.NewLimiter(r, b)
}
