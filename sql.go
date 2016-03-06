package main

import (
	"database/sql"
	"net/http"

	"github.com/NorgannasAddOns/go-uuid"
)

func lookupUserById(id string) (firstname, lastname, email, hashed string, err error) {
	err = db.QueryRow("SELECT FirstName, LastName, Email, Password FROM User WHERE ID=?", id).
		Scan(&firstname, &lastname, &email, &hashed)
	switch {
	case err == sql.ErrNoRows:
		err = &apiError{http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)}
	}

	return
}

func lookupUserByEmail(email string) (firstname, lastname, id, hashed string, err error) {
	err = db.QueryRow("SELECT Firstname, Lastname, ID, Password FROM User WHERE Email=?", email).
		Scan(&firstname, &lastname, &id, &hashed)
	switch {
	case err == sql.ErrNoRows:
		err = &apiError{http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)}
	}

	return
}

func updatePassword(id, password string) error {
	stmt, err := db.Prepare("UPDATE Users SET Password=? WHERE ID=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(password, id)
	return err
}

func createUser(email, firstname, lastname, password string) (string, error) {
	hashed := encodePassword(password)
	id := uuid.New("U")

	stmt, err := db.Prepare("INSERT INTO User (ID, FirstName, LastName, Email, Password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(id, firstname, lastname, email, hashed)
	if err != nil {
		return "", err
	}

	return id, nil
}
