package jsonKit

import (
	"log"
	"testing"

	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
)

func TestMarshalToFile(t *testing.T) {
	log.Printf("os: %s", osKit.OS)
	log.Printf("arch: %s", osKit.ARCH)
	log.Printf("json library: %s", library)

	m := map[string]interface{}{
		"a": 1,
		"b": []string{"0", "1", "2"},
	}
	if err := MarshalToFile(m, "_test.json", 0666); err != nil {
		panic(err)
	}
	if err := MarshalToFileWithAPI(GetStdApi(), m, "_test1.json", 0666); err != nil {
		panic(err)
	}
}
