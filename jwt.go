package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/freman/authserve/token"
)

func generateToken(c claims) *jwt.Token {
	t := jwt.New(jwt.SigningMethodRS256)
	t.Header["typ"] = "JWT"

	for name, value := range c {
		t.Claims[name] = value
	}

	t.Claims["exp"] = time.Now().Add(config.JWT.Expiry.Duration).Unix()
	t.Claims["nbf"] = time.Now().Add(-config.JWT.NotBefore.Duration).Unix()
	t.Claims["iat"] = time.Now().Unix()

	return t
}

func verifyToken(t string) (map[string]interface{}, error) {
	claims, err := token.Verify(t, config.JWT.Public.Bytes())
	if err != nil {
		switch err {
		case token.ErrMissingClaims:
			return nil, &apiError{http.StatusUnauthorized, "Invalid JWT - Missing claims"}
		case token.ErrExpired:
			return nil, &apiError{http.StatusUnauthorized, "Token has expired"}
		case token.ErrNotYetValid:
			return nil, &apiError{http.StatusUnauthorized, "Token is not yet valid"}
		}
		return nil, &apiError{http.StatusUnauthorized, "Invalid token supplied - " + err.Error()}
	}

	return claims, nil
}
