package reqKit

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/richelieu042/chimera/v3/src/netKit"
)

// POST请求，"Content-Type"为"application/x-www-form-urlencoded; charset=utf-8"
func TestPost(t *testing.T) {
	go func() {
		port := 8001

		engine := gin.Default()
		engine.POST("/test", func(ctx *gin.Context) {
			name := ctx.PostForm("name")
			age := ctx.PostForm("age")
			ctx.String(200, fmt.Sprintf("Hello %s(%s)", name, age))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	// 设置请求数据
	data := url.Values{}
	data.Set("name", "张三")
	data.Set("age", "30")

	resp := Post(context.TODO(), "http://127.0.0.1:8001/test", ContentTypeForm, data.Encode())
	if resp.Err != nil {
		panic(resp.Err)
	}
	fmt.Println("Response Status:", resp.Status)
	fmt.Println(resp.ToString())
}

// POST请求，"Content-Type"为"application/json; charset=utf-8"
func TestPost1(t *testing.T) {
	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	go func() {
		port := 8001

		engine := gin.Default()
		engine.POST("/test", func(ctx *gin.Context) {
			p := &person{}
			if err := ctx.Bind(p); err != nil {
				ctx.String(500, err.Error())
				return
			}
			ctx.String(200, fmt.Sprintf("Hello %s(%d)", p.Name, p.Age))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	// 创建请求客户端
	client := req.C()

	// 设置请求数据
	p := &person{
		Name: "李四",
		Age:  40,
	}

	// 发送POST请求
	resp, err := client.R().
		//SetContentType("application/json; charset=utf-8").
		SetBody(p).
		Post("http://127.0.0.1:8001/test")

	if err != nil {
		panic(err)
	}

	// 处理响应
	fmt.Println("Response Status:", resp.Status)
	fmt.Println(resp.ToString())
}
