package tomlKit

import (
	"fmt"
	"testing"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

func TestUnmarshal(t *testing.T) {
	type Config struct {
		Age        int
		Cats       []string
		Pi         float64
		Perfection []int
		DOB        time.Time
	}

	text := `
Age = 25
Cats = [ "Cauchy", "Plato" ]
Pi = 3.14
Perfection = [ 6, 28, 496, 8128 ]
DOB = 1987-07-05T05:45:00Z
`
	var conf Config
	err := toml.Unmarshal([]byte(text), &conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonKit.MarshalIndentToString(conf, "", "  "))
}
