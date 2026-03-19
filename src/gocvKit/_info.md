## 参考

- [GitHub](https://github.com/hybridgroup/gocv)

## 关闭 gocv.Mat 实例

var img gocv.Mat 只是声明了零值 Mat（相当于 gocv.Mat{}），没有调用任何分配 C 内存的函数，所以没有需要释放的资源。

只有以下情况才需要 .Close()：

- gocv.NewMat()
- gocv.NewMatWithSize(...)
- gocv.IMRead(...)
- gocv.CvtColor 的输出 Mat（你代码里 img = gocv.NewMat() 那种）
- 其他会返回/填充新 Mat 的函数

### case: 零值（gocv.Mat{}）调用 Close()

一切正常，会返回 nil（并不会发生 panic 等）。
但不推荐这么干。
