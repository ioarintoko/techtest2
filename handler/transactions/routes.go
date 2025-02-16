package transactions

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Route(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		Insert(db, w, r)

	case http.MethodPut:
		Update(db, w, r)

	case http.MethodGet:
		Gets(db, w, r)

	default:
		fmt.Println("Wrong Method Admin")
	}
}
