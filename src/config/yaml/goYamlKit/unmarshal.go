package goYamlKit

import "github.com/goccy/go-yaml"

func Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
