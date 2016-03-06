package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func authorize(w http.ResponseWriter, r *http.Request, s securedHandler) error {
	token := r.Header.Get("authorization")
	if token == "" {
		return &apiError{http.StatusUnauthorized, "Missing authorization header"}
	}

	claims, err := verifyToken(token)
	if err != nil {
		return err
	}

	return s(claims, w, r)
}

func jsonBody(r *http.Request, v interface{}) error {
	contentType := r.Header.Get("content-type")
	if !strings.HasPrefix(contentType, "application/json") {
		return &apiError{http.StatusUnsupportedMediaType, fmt.Sprintf("%s (expected application/json got %s)", http.StatusText(http.StatusUnsupportedMediaType), contentType)}
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		return &apiError{http.StatusBadRequest, "Invalid JSON provided"}
	}
	return nil
}

func authenticate(w http.ResponseWriter, r *http.Request) error {
	req := authenticateRequest{}
	if err := jsonBody(r, &req); err != nil {
		return err
	}

	if len(req.Email) < 3 {
		return &apiError{http.StatusBadRequest, "Email is not long enough"}
	}

	if len(req.Password) < 2 {
		return &apiError{http.StatusBadRequest, "Password is not long enough"}
	}

	firstname, lastname, id, hashed, err := lookupUserByEmail(req.Email)
	if err != nil {
		return err
	}

	err = verifyPassword(hashed, req.Password)
	if err != nil {
		return &apiError{http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)}
	}

	token := generateToken(claims{
		"id":        id,
		"email":     req.Email,
		"firstname": firstname,
		"lastname":  lastname,
	})

	signedToken, err := token.SignedString(config.JWT.Private.Bytes())
	if err != nil {
		return err
	}

	js, err := json.Marshal(&tokenResponse{signedToken})
	if err != nil {
		return err
	}
	w.Write(js)

	return nil
}

func refreshToken(c claims, w http.ResponseWriter, r *http.Request) error {
	token := generateToken(c)

	signedToken, err := token.SignedString(config.JWT.Private.Bytes())
	if err != nil {
		return err
	}

	js, err := json.Marshal(&tokenResponse{signedToken})
	if err != nil {
		return err
	}
	w.Write(js)

	return nil
}

func reloadUser(c claims, w http.ResponseWriter, r *http.Request) error {
	id := c["id"].(string)

	var err error
	c["firstname"], c["lastname"], c["email"], _, err = lookupUserById(id)
	if err != nil {
		return err
	}

	token := generateToken(c)

	signedToken, err := token.SignedString(config.JWT.Private.Bytes())
	if err != nil {
		return err
	}

	js, err := json.Marshal(&tokenResponse{signedToken})
	if err != nil {
		return err
	}
	w.Write(js)

	return nil
}

func changePassword(c claims, w http.ResponseWriter, r *http.Request) error {
	req := passwordRequest{}
	if err := jsonBody(r, &req); err != nil {
		return err
	}

	if len(req.Password) < 6 {
		return &apiError{http.StatusPreconditionFailed, "Password is required to be a minimum of 6 characters"}
	}

	id := c["id"].(string)

	if err := updatePassword(id, req.Password); err != nil {
		return err
	}

	// Cheating :D
	return &apiError{http.StatusOK, "Password changed"}
}

func newUser(w http.ResponseWriter, r *http.Request) error {
	req := createUserRequest{}
	if err := jsonBody(r, &req); err != nil {
		return err
	}

	if len(req.Password) < 6 {
		return &apiError{http.StatusPreconditionFailed, "Password is required to be a minimum of 6 characters"}
	}

	id, err := createUser(req.Email, req.Firstname, req.Lastname, req.Password)
	if err != nil {
		return err
	}

	token := generateToken(claims{
		"id":        id,
		"email":     req.Email,
		"firstname": req.Firstname,
		"lastname":  req.Lastname,
	})

	signedToken, err := token.SignedString(config.JWT.Private.Bytes())
	if err != nil {
		return err
	}

	js, err := json.Marshal(&tokenResponse{signedToken})
	if err != nil {
		return err
	}
	w.Write(js)

	return nil
}

func swagger(w http.ResponseWriter, r *http.Request) error {
	w.Write(swaggerDefinition)
	return nil
}
