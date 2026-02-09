package propertiesKit

import (
	"github.com/magiconair/properties"
	"github.com/mitchellh/mapstructure"
	"github.com/richelieu042/chimera/v3/src/core/mapKit"
)

var (
	Load func(buf []byte, enc properties.Encoding) (*properties.Properties, error) = properties.Load

	LoadString func(s string) (*properties.Properties, error) = properties.LoadString

	LoadFile  func(filename string, enc properties.Encoding) (*properties.Properties, error)                        = properties.LoadFile
	LoadFiles func(filenames []string, enc properties.Encoding, ignoreMissing bool) (*properties.Properties, error) = properties.LoadFiles
)

func Unmarshal(content []byte, ptr interface{}) error {
	return unmarshal(content, properties.UTF8, ptr)
}

func UnmarshalWithIso88591(content []byte, ptr interface{}) error {
	return unmarshal(content, properties.ISO_8859_1, ptr)
}

// unmarshal
/*
@param ptr 可以是:	var obj interface{}
					unmarshal(content, enc, &obj)
*/
func unmarshal(content []byte, enc properties.Encoding, ptr interface{}) error {
	p, err := Load(content, enc)
	if err != nil {
		return err
	}

	/*
		map[string]string => map[string]interface{}
		加这个主要是为了 i18nKit读取.properties文件 的需要
	*/
	var m map[string]interface{} = mapKit.MapValues(p.Map(), func(value string, key string) interface{} {
		return value
	})

	return mapstructure.Decode(m, ptr)
}
