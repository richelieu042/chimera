package jsonKit

var (
	library string

	// defaultApi 默认的API
	defaultApi API = nil

	// stdApi 标准的API（会对map的keys排序）
	stdApi API = nil
)

type (
	API interface {
		Marshal(v any) ([]byte, error)

		MarshalIndent(v any, prefix, indent string) ([]byte, error)

		MarshalToString(v any) (string, error)

		Unmarshal(data []byte, v any) error

		UnmarshalFromString(str string, v any) error
	}
)

func GetLibrary() string {
	return library
}

func GetDefaultApi() API {
	return defaultApi
}

func GetStdApi() API {
	return stdApi
}
