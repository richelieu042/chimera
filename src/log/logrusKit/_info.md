## 参考

- notes/Golang/日志 log/logrus.wps

## 简介

Logrus是一个结构化、分级、可扩展和兼容标准库log包的日志库，由Sirupsen开发和维护。
Logrus提供了一个简单而强大的API，可以用于记录不同级别和格式的日志信息。

Logrus也提供了一些有用的特性，如：

* 可以自定义日志级别、输出目标、格式器（JSON或文本）、时间戳等；
* 可以使用字段（Fields）来添加结构化的上下文信息，如键值对；
* 可以使用钩子（Hooks）来在每次写入日志时执行一些操作，如发送邮件、写入数据库等；
* 可以使用条目（Entry）来记录带有字段的日志信息，或者使用WithFields、WithTime、WithError等方法来创建带有字段的条目；
* 可以使用日志级别函数（如Info、Warn、Error等）来记录不同级别的日志信息，或者使用Log或Print等方法来记录默认级别的日志信息。

## logrus库

- 默认的日志级别: INFO
- 默认的Out属性: os.Stderr

## exit handler（退出函数）

e.g.

```go
package main

import (
    "github.com/sirupsen/logrus"
)

func main() {
    // 程序正常退出不会触发 logrus.RegisterExitHandle() 注册的exit handlers，因此加上此defer语句.
    defer logrus.Fatal("main() ends")

	logrus.RegisterExitHandler(func() {
		logrus.Info(111)
	})
}
```

## 有用的实践

#### 特定级别的日志输出，也同时输出到别的地方（比如: 控制台 || 文件）

详见: hook_lfshook.go.
