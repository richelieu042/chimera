## Settings => Build Tags => Custom tags

- sonic: 是否想用sonic库？（但是并不一定最终使用sonic库，还要看其他条件）
- avx: CPU是否支持AVX指令集？（针对amd64架构）

## 想要获取指定字段的值，但又不想定义 struct 去 Unmarshal 来获取字段的值

结合gjson轻松获取json字段
https://req.cool/zh/docs/tutorial/handle-response/#%e7%bb%93%e5%90%88-gjson-%e8%bd%bb%e6%9d%be%e8%8e%b7%e5%8f%96-json-%e5%ad%97%e6%ae%b5

如果返回的响应体是 json 格式，想要获取指定字段的值，但又不想定义 struct 去 Unmarshal 来获取字段的值，可以结合 gjson
来获取指定字段的值。

## !!!: 缺陷

#### 反序列化为map，key的类型不能是interface{}（json-iterator/go返回error）

```go
package main

import (
	"fmt"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

func main() {
	m := map[interface{}]interface{}{
		"0": 3.1415926,
	}

	data, err := jsonKit.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	m1 := map[interface{}]interface{}{}
	if err := jsonKit.Unmarshal(data, &m1); err != nil {
		panic(err) // panic: unsupported map key type: interface {}
	}
	fmt.Println(m1)
}
```

