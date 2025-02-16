package authentications

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

func Route(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if strings.Split(r.URL.Path, "/")[2] == "login" {
			Login(db, w, r)
		} else {
			Logout(w, r)
		}

	default:
		fmt.Println("Wrong Auth Method")
	}
}
