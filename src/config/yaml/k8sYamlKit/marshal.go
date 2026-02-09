package k8sYamlKit

import (
	"os"

	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"sigs.k8s.io/yaml"
)

var Marshal func(o interface{}) ([]byte, error) = yaml.Marshal

func MarshalToString(in interface{}) (string, error) {
	data, err := Marshal(in)
	return string(data), err
}

func MarshalToFile(in interface{}, filePath string, perm os.FileMode) error {
	data, err := Marshal(in)
	if err != nil {
		return err
	}
	return fileKit.WriteToFile(filePath, data)
}
