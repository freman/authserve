package main

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

var ErrPasswordIncorrect = errors.New("Password incorrect")

func encodePassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(hashed)
}

func verifyPassword(hashed, password string) error {
	// Bcrypt actually returns a bunch of different errors but they shouldn't happen due to upstream filtering
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err == nil {
		return nil
	}
	return ErrPasswordIncorrect
}
