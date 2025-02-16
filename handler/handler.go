package handler

import (
	"database/sql"
	users "techtest2/handler/Users"
	"techtest2/handler/authentications"

	"fmt"
	"net/http"
	"strings"
)

var DB *sql.DB

func RegisDB(db *sql.DB) {
	DB = db
}

const (
	login    = "login"
	logout   = "logout"
	profile  = "profile"
	register = "register"
)

func API(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Header", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	url := r.URL.Path
	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")

	fmt.Println("token center", w.Header()["Set-Cookie"])

	switch dataURL[2] {
	case login:
		authentications.Auth(DB, w, r)

	case logout:
		authentications.Auth(DB, w, r)

	case register:
		users.Users(DB, w, r)

	case profile:
		users.Users(DB, w, r)

	default:
		fmt.Println("Wrong Path")

	}
}
