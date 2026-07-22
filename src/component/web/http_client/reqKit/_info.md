## Deprecated!!!

- 国人作者，很久不更新了
- // 当前最新版 v3.57.0 使用的还是 github.com/quic-go/quic-go v0.57.1，这会和 github.com/gin-gonic/gin v1.12.0 冲突，因此被移出
  chimera.

## 参考

- [imroc/req](https://github.com/imroc/req)
- notes/Golang/WEB/req - http客户端.wps

## Content-Type

imcro/req 发送的POST请求，默认的 "Content-Type" 是 "application/json; charset=utf-8" .

## 自动重试 (retry)

默认情况下，

- 不retry
- retry interval为100ms


