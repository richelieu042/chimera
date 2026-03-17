## 注意点

- 为了防止import cycle，不要引用 console 包.

## 参考

- notes/Golang/log（日志）/zap.docx
- [github 21.2k](https://github.com/uber-go/zap)
- [Go 语言之 zap 日志库简单使用](https://zhuanlan.zhihu.com/p/637747131)

## zapcore.WriteSyncer接口的实例

- zapcore.AddSync 的返回值
- zapcore.Lock 的返回值

## TODOs

- wrap.go中的方法，要针对情况: 并发情况下，通过已经关闭了logger继续输出日志. 
   
 
