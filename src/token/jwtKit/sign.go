package jwtKit

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/mapKit"
	"github.com/richelieu042/chimera/v3/src/crypto/caesarKit"
)

// Sign 生成JWT字符串.
/*
@param key		密钥（secret）
@param method 	e.g. jwt.SigningMethodHS256 || jwt.SigningMethodHS384 || jwt.SigningMethodHS512
*/
func Sign(key []byte, method jwt.SigningMethod, claims jwt.MapClaims, options ...jwt.TokenOption) (string, error) {
	if err := interfaceKit.AssertNotNil(key, "key"); err != nil {
		return "", err
	}
	if err := interfaceKit.AssertNotNil(method, "method"); err != nil {
		return "", err
	}
	if err := mapKit.AssertNotEmpty(claims, "claims"); err != nil {
		return "", err
	}

	/*
		Create a new token object, specifying signing method and the claims
		you would like it to contain.
	*/
	token := jwt.NewWithClaims(method, claims, options...)

	/*
		Sign and get the complete encoded token as a string using the secret
	*/
	return token.SignedString(key)
}

func SignWithCaesar(caesarShift int, key []byte, method jwt.SigningMethod, claims jwt.MapClaims, options ...jwt.TokenOption) (cipherText string, err error) {
	var tokenString string
	tokenString, err = Sign(key, method, claims, options...)
	if err != nil {
		return
	}

	cipherText = caesarKit.Encrypt(tokenString, caesarShift)
	return
}
