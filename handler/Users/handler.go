package users

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

	if lastIndex == "register" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()
		var user model.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		datauser, err := user.Insert(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(datauser)
	}
}

func Update(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dataUrl := strings.Split(url, "/")
	lastIndex := dataUrl[len(dataUrl)-1]

	if lastIndex == "profile" {
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
		datauser := model.User{IDUser: iduser}
		dataUpdate, err := datauser.Update(db, jsonMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dataUpdate)
	}
}
