## natefinch/lumberjack 库已长期未更新

可以考虑使用 [DeRuina/timberjack](https://github.com/DeRuina/timberjack) 作为替代.

## bytes.Buffer 结构体

Golang 语言标准库 bytes 包怎么使用？
https://mp.weixin.qq.com/s?__biz=MzA4Mjc1NTMyOQ==&mid=2247484410&idx=1&sn=e50b07f7adf7ac0c5f6d0496509d34fa

一个可读写、可变大小的字节缓冲区. bytes.NewBuffer 传参可以为nil.

- 实现了: io.Reader、io.Writer
- 未实现: io.Seeker、io.Closer

## bytes.Reader 结构体

Golang 语言标准库 bytes 包怎么使用？
https://mp.weixin.qq.com/s?__biz=MzA4Mjc1NTMyOQ==&mid=2247484410&idx=1&sn=e50b07f7adf7ac0c5f6d0496509d34fa

可以读取 []byte。 与 Buffer 可读写不同，Reader 是只读和支持查找。

- 实现了: io.Reader、io.Seeker
- 未实现: io.Writer、io.Closer

## strings.Reader 结构体

- 实现了: io.Reader、io.Seeker
- 未实现: io.Writer、io.Closer

## bufio.Reader 结构体

- 实现了: io.Reader
- 未实现: io.Writer、io.Seeker、io.Closer

## bufio.Writer 结构体

- 实现了: io.Writer
- 未实现: io.Reader、io.Seeker、io.Closer

## bufio.ReadWriter 结构体

- 实现了: io.Reader、io.Writer
- 未实现: io.Seeker、io.Closer

## os.File 结构体

- 实现了: io.Reader、io.Writer、io.Seeker、io.Closer


