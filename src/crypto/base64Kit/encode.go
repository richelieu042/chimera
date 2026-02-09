package base64Kit

import (
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

func Encode(src []byte, options ...Base64Option) []byte {
	opts := loadOptions(options...)
	return opts.Encode(src)
}

func EncodeToString(src []byte, options ...Base64Option) string {
	opts := loadOptions(options...)
	return opts.EncodeToString(src)
}

// EncodeStringToString (拓展) string => base64 string
func EncodeStringToString(s string, options ...Base64Option) string {
	return EncodeToString([]byte(s), options...)
}

// EncodeFile （拓展）file => []byte
func EncodeFile(path string, options ...Base64Option) ([]byte, error) {
	data, err := fileKit.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return Encode(data, options...), nil
}

// EncodeFileToString （拓展）file => string
func EncodeFileToString(path string, options ...Base64Option) (string, error) {
	data, err := fileKit.ReadFile(path)
	if err != nil {
		return "", err
	}

	return EncodeToString(data, options...), nil
}
