package viperKit

import (
	"github.com/go-viper/encoding/ini"
	"github.com/go-viper/encoding/javaproperties"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/ioKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/spf13/viper"
)

// newViper 创建一个增强了 codec 支持的 viper 实例（在内置 json/toml/yaml/dotenv 基础上，额外支持 properties、ini）.
/*
viper库从 v1.20.0 开始支持，移除了对 HCL、Java properties、INI  的支持，详见: https://github.com/spf13/viper/pull/1870
*/
func newViper() *viper.Viper {
	codecRegistry := viper.NewCodecRegistry()

	// 支持: properties
	pc := &javaproperties.Codec{}
	_ = codecRegistry.RegisterCodec("properties", pc)
	_ = codecRegistry.RegisterCodec("props", pc)
	_ = codecRegistry.RegisterCodec("prop", pc)

	// 支持: ini
	ic := &ini.Codec{}
	_ = codecRegistry.RegisterCodec("ini", ic)

	return viper.NewWithOptions(viper.WithCodecRegistry(codecRegistry))
}

func Read(data []byte, configType string, defaultMap map[string]interface{}) (*viper.Viper, error) {
	if err := interfaceKit.AssertNotNil(data, "data"); err != nil {
		return nil, err
	}
	if err := strKit.AssertNotBlank(configType, "configType"); err != nil {
		return nil, err
	}
	configType = PolyfillContentType(configType)

	v := newViper()
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

	v := newViper()
	for key, value := range defaultMap {
		v.SetDefault(key, value)
	}
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}
