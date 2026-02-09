package pushKit

import "github.com/richelieu042/chimera/v3/src/idKit"

func DefaultIdGenerator() func() (string, error) {
	return func() (string, error) {
		return idKit.NewXid(), nil
	}
}
