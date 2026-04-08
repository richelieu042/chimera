package main

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/config/viperKit"
)

func main() {
	path := "/Users/richelieu/Downloads/message.properties"

	m := map[string]string{}
	v, err := viperKit.UnmarshalFromFile(path, nil, &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
