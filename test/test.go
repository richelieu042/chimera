package main

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/gocvKit"
	"gocv.io/x/gocv"
)

func main() {
	mat, err := gocvKit.DecodeFromPath("111.png")
	defer mat.Close()
	if err != nil {
		panic(err)
	}

	// 调用时
	data, _ := MatToBytesWrong(mat, ".png")
	fmt.Println(data) // 可能输出乱码，或 panic，行为不确定
}

// ❌ 错误写法 — buf.Close() 后内存被释放，data 变成野指针
func MatToBytesWrong(mat gocv.Mat, format string) ([]byte, error) {
	buf, err := gocv.IMEncode(gocv.FileExt(format), mat)
	if err != nil {
		return nil, err
	}
	defer buf.Close() // Close 后下面的 data 底层内存已释放

	data := buf.GetBytes() // 直接返回引用，不安全
	return data, nil
}
