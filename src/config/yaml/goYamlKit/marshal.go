package goYamlKit

import (
	"github.com/goccy/go-yaml"
)

func Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}
