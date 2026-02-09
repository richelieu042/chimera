package viperKit

import (
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/ioKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/spf13/viper"
)

func Read(data []byte, configType string, defaultMap map[string]interface{}) (*viper.Viper, error) {
	if err := interfaceKit.AssertNotNil(data, "data"); err != nil {
		return nil, err
	}
	if err := strKit.AssertNotBlank(configType, "configType"); err != nil {
		return nil, err
	}
	configType = PolyfillContentType(configType)

	v := viper.New()
	for key, value := range defaultMap {
		v.SetDefault(key, value)
	}
	v.SetConfigType(configType)
	if err := v.ReadConfig(ioKit.NewReader(data)); err != nil {
		return nil, err
	}
	return v, nil
}

func ReadFile(filePath string, defaultMap map[string]interface{}) (*viper.Viper, error) {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}

	v := viper.New()
	for key, value := range defaultMap {
		v.SetDefault(key, value)
	}
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}
