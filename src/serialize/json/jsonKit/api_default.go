//go:build !(go1.18 && amd64 && sonic && avx) && !(go1.20 && arm64 && sonic)

package jsonKit

import "encoding/json"

type defaultImpl struct {
}

func (i *defaultImpl) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (i *defaultImpl) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (i *defaultImpl) MarshalToString(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}

func (i *defaultImpl) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (i *defaultImpl) UnmarshalFromString(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}

func init() {
	library = "encoding/json"

	tmp := &defaultImpl{}
	defaultApi = tmp
	stdApi = tmp
}
