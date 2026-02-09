package kratosKit

import (
	"encoding/json"
	"reflect"

	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"

	"github.com/go-kratos/kratos/v2/encoding"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// RegisterDefaultCodec
/*
参考: 序列化 https://go-kratos.dev/docs/component/encoding
*/
func RegisterDefaultCodec() {
	var p encoding.Codec = &codec{}
	encoding.RegisterCodec(p)
}

var (
	// marshalOptions is a configurable JSON format marshaller.
	marshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	// unmarshalOptions is a configurable JSON format parser.
	unmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

type codec struct{}

func (*codec) Name() string {
	return "json"
}

func (*codec) Marshal(v interface{}) ([]byte, error) {
	switch m := v.(type) {
	case json.Marshaler:
		return m.MarshalJSON()
	case proto.Message:
		return marshalOptions.Marshal(m)
	default:
		return jsonKit.Marshal(m)
	}
}

func (*codec) Unmarshal(data []byte, v interface{}) error {
	switch m := v.(type) {
	case json.Unmarshaler:
		return m.UnmarshalJSON(data)
	case proto.Message:
		return unmarshalOptions.Unmarshal(data, m)
	default:
		rv := reflect.ValueOf(v)
		for rv := rv; rv.Kind() == reflect.Ptr; {
			if rv.IsNil() {
				rv.Set(reflect.New(rv.Type().Elem()))
			}
			rv = rv.Elem()
		}
		if m, ok := reflect.Indirect(rv).Interface().(proto.Message); ok {
			return unmarshalOptions.Unmarshal(data, m)
		}
		return jsonKit.Unmarshal(data, m)
	}
}
