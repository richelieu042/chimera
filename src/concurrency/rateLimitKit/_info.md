## 参考

- https://claude.ai/share/2fe54c29-a922-4f24-9400-a35bc8b368dd

## go.uber.org/ratelimit

### slack（松弛）

https://claude.ai/share/ffed7673-aebd-4f59-a00c-d387e545ff3b

uber-go/ratelimit 默认行为是**有 slack（松弛）**的：如果你有一段时间没有调用 Take()，limiter 会"积累"最多 10
个令牌的空闲时间，下次可以连续快速取。

WithoutSlack 禁用这个积累，严格保证每两次 Take() 之间的间隔不小于 1/rate，无论之前空闲了多久。