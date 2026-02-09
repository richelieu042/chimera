package jwtKit

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

func IsTokenExpiredError(err error) bool {
	return strKit.Contains(err.Error(), jwt.ErrTokenExpired.Error())
}
