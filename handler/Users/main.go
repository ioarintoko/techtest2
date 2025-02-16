package users

import (
	"database/sql"
	"net/http"
)

func Users(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Route(db, w, r)
}
