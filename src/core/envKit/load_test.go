package envKit

import (
	"fmt"
	"testing"

	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

func TestReadFromFile(t *testing.T) {
	m, err := ReadFromFile("./_test.yaml", "./_test.env")
	if err != nil {
		panic(err)
	}
	json, err := jsonKit.MarshalIndentToString(m, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(json)
	if err := SetEnvs(m); err != nil {
		panic(err)
	}
	fmt.Println(GetEnv("c"))
}
