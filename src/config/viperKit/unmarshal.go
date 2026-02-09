// Package viperKit
/*
支持的配置文件格式（详见viper.go）:
"yaml", "yml", "json", "toml", "hcl", "tfvars",
"dotenv", "env",
"properties", "props", "prop",
"ini"
*/
package viperKit

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/ptrKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/spf13/viper"
)

// Unmarshal 读取配置文本内容，并反序列化.
/*
@param configType 	配置文件的类型（不区分大小写，详见viper.go）："yaml", "yml", "json", "toml", "hcl", "tfvars", "dotenv", "env", "properties", "props", "prop", "ini"
@param defaultMap 	默认值，可以为nil
@param ptr			指针，且不能为nil
*/
func Unmarshal(content []byte, configType string, defaultMap map[string]interface{}, ptr interface{}) (*viper.Viper, error) {
	if err := ptrKit.AssertNotNilAndIsPointer(ptr); err != nil {
		return nil, err
	}

	v, err := Read(content, configType, defaultMap)
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(ptr, func(dc *mapstructure.DecoderConfig) {
		// 如果指针ptr对应的类型是个子类的话，需要如此进行设置，否则父类的属性都会是对应类型的零值！
		dc.Squash = true
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}

func UnmarshalFromData(data *Data, defaultMap map[string]interface{}, ptr interface{}) (*viper.Viper, error) {
	if err := interfaceKit.AssertNotNil(data, "data"); err != nil {
		return nil, err
	}

	return Unmarshal(data.Content, data.Type, defaultMap, ptr)
}

// UnmarshalFromFile 读取配置文件，并反序列化.
/*
PS:
(1) 配置文件 和 defaultMap 中，key首字母的大小写无所谓，都支持；
(2) 支持配置文件的格式：JSON, TOML, HCL, .env, .yaml, .properties.

@param filePath		配置文件的路径（绝对路径 和 相对路径 都支持），内部会判断文件是否存在
@param defaultMap 	（可以为nil） 默认值；key如果有多层的话，用"."分隔，e.g. "WoService.LowerLimit"
@param ptr			（不能为nil） 指针
*/
func UnmarshalFromFile(filePath string, defaultMap map[string]interface{}, ptr interface{}) (*viper.Viper, error) {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}
	if err := ptrKit.AssertNotNilAndIsPointer(ptr); err != nil {
		return nil, err
	}

	v, err := ReadFile(filePath, defaultMap)
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(ptr, func(dc *mapstructure.DecoderConfig) {
		// 如果指针ptr对应的类型是个子类的话，需要如此进行设置，否则父类的属性都会是对应类型的零值！
		dc.Squash = true
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}
