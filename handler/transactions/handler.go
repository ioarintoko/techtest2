package transactions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"techtest2/model"
	"techtest2/tokenize"

	"github.com/golang-jwt/jwt/v4"
)

func Insert(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dataUrl := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataUrl[len(dataUrl)-1]

	if lastIndex == "transaction" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()
		var transaction model.Transaction
		err = json.Unmarshal(body, &transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		datatr, err := transaction.Insert(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(datatr)
	}
}

func Update(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dataUrl := strings.Split(url, "/")
	lastIndex := dataUrl[len(dataUrl)-1]
	endpoint := dataUrl[len(dataUrl)-2]

	if endpoint == "pay" || endpoint == "topup" || endpoint == "transfer" || endpoint == "transaction" {
		// Extract JWT Token
		token, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		var cookie jwt.MapClaims
		if token != nil {
			cookie = tokenize.Decode(w, r)
		}

		iduser, ok := cookie["iduser"].(string)
		if !ok || iduser == "" {
			http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
			return
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Parse JSON body
		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(body, &jsonMap)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// Update user
		datatr := model.Transaction{IDTransaction: lastIndex}
		dataUpdate, err := datatr.Update(db, jsonMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dataUpdate)
	}
}

func Gets(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var cookie jwt.MapClaims
	token, err := r.Cookie("token")
	if err != nil {
		fmt.Println("error handler gets", err)
	}

	if token != nil {
		cookie = tokenize.Decode(w, r)
	}

	var iduser string
	if cookie != nil {
		iduser = cookie["iduser"].(string)
	} else {
		iduser = ""
	}

	postDatas, err := model.GetsTransaction(db, iduser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jsonData, err := json.Marshal(postDatas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(jsonData)
}
