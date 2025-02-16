package transactions

import (
	"database/sql"
	"net/http"
)

func Transactions(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Route(db, w, r)
}
