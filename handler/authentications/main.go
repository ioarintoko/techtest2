package authentications

import (
	"database/sql"
	"net/http"
)

func Auth(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Route(db, w, r)
}
