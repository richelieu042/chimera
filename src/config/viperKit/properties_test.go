package viperKit

import (
	"fmt"
	"testing"

	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

// TestLoadProperties 读取"_test.properties".
func TestLoadProperties(t *testing.T) {
	type redis struct {
		Cluster []string `json:"cluster"`
	}
	type config struct {
		Redis redis `json:"redis"`
	}

	path := "_test.properties"
	c := &config{}
	_, err := UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}

	str, _ := jsonKit.MarshalIndentToString(c, "", "    ")
	fmt.Println(str)
}
