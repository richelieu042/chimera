package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"regexp"
)

func main() {
	path := "/Users/richelieu/Downloads/剑啸灵霄(1-380章).txt"
	content, err := fileKit.ReadFileToString(path)
	if err != nil {
		panic(err)
	}

	re, err := regexp.Compile(`(\d+)、(.+)`)
	if err != nil {
		panic(err)
	}
	content1 := re.ReplaceAllStringFunc(content, func(str string) string {
		s := re.FindStringSubmatch(str)
		s1 := fmt.Sprintf("第%s章、%s", s[1], s[2])

		fmt.Println(s[1])

		return s1
	})
	//fmt.Println(content1)

	if err := fileKit.WriteStringToFile("/Users/richelieu/Downloads/剑啸灵霄.txt", content1, false); err != nil {
		panic(err)
	}

	//engine := gin.Default()
	//
	//engine.Any("/test", func(ctx *gin.Context) {
	//	console.Info("---")
	//	s := []string{
	//		"X-Forwarded-For",
	//		"X-Real-IP",
	//		"Forwarded",
	//		"Via",
	//		"CF-Connecting-IP",
	//		"True-Client-IP",
	//	}
	//	for _, key := range s {
	//		value := ctx.Request.Header.Get(key)
	//		console.Infof("[HEADER] key: %s, value: %s", key, value)
	//	}
	//
	//	console.Infof("ClientIP: %s", ctx.ClientIP())
	//	console.Infof("RemoteIP: %s", ctx.RemoteIP())
	//	console.Info("===")
	//
	//	ctx.String(200, "Hello world!")
	//})
	//
	//if err := engine.Run(":8001"); err != nil {
	//	panic(err)
	//}
}
