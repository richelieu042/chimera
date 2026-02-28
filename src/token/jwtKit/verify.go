package jwtKit

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/crypto/caesarKit"
)

// Verify 验证JWT字符串.
/*
PS: 如果 token 过期（根据"exp"，有的话），会返回 error（可以通过 IsTokenExpiredError 判断）.

@param keyFunc 此函数的第一个返回值为密钥(secret)
	e.g.
	func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return key, nil
	}
*/
func Verify(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (jwt.MapClaims, error) {
	if err := strKit.AssertNotEmpty(tokenString, "tokenString"); err != nil {
		return nil, err
	}
	if len(strKit.Split(tokenString, ".")) != 3 {
		return nil, errorKit.Newf("tokenString(%s) is invalid", tokenString)
	}
	if err := interfaceKit.AssertNotNil(keyFunc, "keyFunc"); err != nil {
		return nil, err
	}

	// Verify takes the token string and a function for looking up the Key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which Key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, keyFunc, options...)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errorKit.Newf("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errorKit.Newf("type(%T) of claims is invalid", token.Claims)
	}
	return claims, nil
}

func VerifyWithCaesar(caesarShift int, cipherText string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (jwt.MapClaims, error) {
	tokenString := caesarKit.Decrypt(cipherText, caesarShift)

	return Verify(tokenString, keyFunc, options...)
}
