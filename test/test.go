package main

import (
	"encoding/json"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"time"
)

type Bean struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
}

func main() {
	b := &Bean{
		ID:        "123",
		UpdatedAt: time.Time{},
	}

	data, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	fmt.Println(jsonKit.GetLibrary())
	fmt.Println(jsonKit.MarshalToString(b))
}
