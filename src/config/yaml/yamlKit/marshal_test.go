package yamlKit

import (
	"fmt"
	"testing"
)

func TestMarshalToString(t *testing.T) {
	{
		s := []string{"foo", "bar"}
		data, err := Marshal(s)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}

	fmt.Println("------")

	{
		m := map[string]any{
			"a": 100,
			"b": 200,
		}
		data, err := Marshal(m)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}
}
