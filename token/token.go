package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrMissingClaims = errors.New("jwt/verify: verification error - missing mandatory claims")
	ErrExpired       = errors.New("jwt/verify: token has expired")
	ErrNotYetValid   = errors.New("jwt/verify: token is not yet valid")
)

func Verify(token string, key []byte) (map[string]interface{}, error) {
	// Workaround numbers becoming floats
	parser := new(jwt.Parser)
	parser.UseJSONNumber = true

	t, err := parser.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	if err == nil && t.Valid {
		var (
			expiry, nbf int64
		)

		_, idOK := t.Claims["id"].(string)
		jExpiry, expiryOK := t.Claims["exp"].(json.Number)
		jNbf, nbfOK := t.Claims["nbf"].(json.Number)

		if expiryOK {
			expiry, err = jExpiry.Int64()
			expiryOK = err == nil
		}
		if nbfOK {
			nbf, err = jNbf.Int64()
			nbfOK = err == nil
		}

		if !(idOK && expiryOK && nbfOK) {
			return nil, ErrMissingClaims
		}

		if expiry < time.Now().Unix() {
			return nil, ErrExpired
		}

		if nbf > time.Now().Unix() {
			return nil, ErrNotYetValid
		}

		return t.Claims, nil
	}

	return nil, err
}
