//go:build !(go1.18 && amd64 && sonic && avx) && !(go1.20 && arm64 && sonic)

package jsonKit

import "encoding/json"

type defaultImpl struct {
}

func (i *defaultImpl) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (i *defaultImpl) MarshalIndent(v any, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (i *defaultImpl) MarshalToString(v any) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}

func (i *defaultImpl) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (i *defaultImpl) UnmarshalFromString(str string, v any) error {
	return json.Unmarshal([]byte(str), v)
}

func init() {
	library = "encoding/json"

	tmp := &defaultImpl{}
	defaultApi = tmp
	stdApi = tmp
}
