package httpClientKit

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://www.baidu.com"

	logrus.Info(0)
	code, data, err := Get(url /*, WithTimeout(time.Second*3)*/)
	logrus.Info(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
	fmt.Println(string(data))
}
