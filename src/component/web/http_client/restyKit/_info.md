## 注意点

- 与标准库net/http（必须手动关闭）不同，Resty 内部已经帮你处理了连接复用和响应体的读取/关闭，不需要手动 resp.Body.Close()。