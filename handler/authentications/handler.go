package authentications

import (
	"database/sql"
	"encoding/json"
	"techtest2/middleware"
	"techtest2/tokenize"

	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	r.Body.Close()

	var login *middleware.Login
	err = json.Unmarshal(body, &login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(login)

	err = login.Login(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		message := "berhasil"
		fmt.Println(message)
		fmt.Println(login)

		tokenize.GetToken(login.IDUser, w, r)
		// jsonData, err := json.Marshal(login)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// }

		// w.Write(jsonData)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenize.Logout(w, r)
}
