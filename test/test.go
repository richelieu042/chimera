package main

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"time"
)

type MyInt[T comparable] = int

type Result[T any] = struct {
	Value T
	Error error
}

type AsyncResult[T any, S ~[]T] = func() (Result[T], S)

type Bean struct {
	ID        string     `json:"id"`
	UpdatedAt *time.Time `json:"updated_at,omitzero"`
}

func main() {
	b := &Bean{
		ID:        "123",
		UpdatedAt: nil,
	}

	// (1) 支持：标准库
	data, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("std:", string(data))

	// (2) 支持：sonic
	str, err := sonic.MarshalString(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("sonic:", str)
}
