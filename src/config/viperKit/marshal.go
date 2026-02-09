package viperKit

import (
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

// MarshalToFile
/*
Deprecated: 不建议用此方法，	(1) 转换为 .yaml 还是用 yamlKit 吧;
							(2) 转换为 .properties 有问题.

@param obj 如果为nil，将返回error
*/
func MarshalToFile(obj interface{}, target string) error {
	if err := interfaceKit.AssertNotNil(obj, "obj"); err != nil {
		return err
	}

	if err := fileKit.AssertNotExistOrIsFile(target); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(target); err != nil {
		return err
	}

	data, err := jsonKit.Marshal(obj)
	if err != nil {
		return err
	}
	v, err := Read(data, "json", nil)
	if err != nil {
		return err
	}
	return v.WriteConfigAs(target)
}
