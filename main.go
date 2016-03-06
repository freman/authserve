package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"

	"github.com/BurntSushi/toml"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type authenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type passwordRequest struct {
	Password string `json:"password"`
}

type createUserRequest struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type claims map[string]interface{}

type securedHandler func(claims, http.ResponseWriter, *http.Request) error

var (
	db     *sql.DB
	config *configuration
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := route(w, r)

	if err != nil {
		if hrr, is := err.(*apiError); is {
			hrr.HTTPError(w)
		} else {
			http.Error(w, "Unexpected internal server error - "+err.Error(), http.StatusInternalServerError)
		}
	}
}

func route(w http.ResponseWriter, r *http.Request) error {
	methodPath := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
	switch methodPath {
	case "POST /authenticate":
		return authenticate(w, r)
	case "POST /password":
		return authorize(w, r, changePassword)
	case "GET /refresh":
		return authorize(w, r, refreshToken)
	case "GET /reload":
		return authorize(w, r, reloadUser)
	case "POST /new":
		return newUser(w, r)
	case "GET /swagger.json":
		return swagger(w, r)
	default:
		return &apiError{http.StatusBadRequest, http.StatusText(http.StatusBadRequest)}
	}
}

func main() {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal("Unable to read configuration file", err)
	}

	var err error
	db, err = sql.Open(config.Database.Driver, config.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handleRequest)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
