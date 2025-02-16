package middleware

import "database/sql"

type Login struct {
	IDUser string `json:"iduser"`
	Phone  string `json:"phone"`
	Pin    string `json:"pin"`
}

func (lgn *Login) Login(db *sql.DB) error {
	query := "SELECT IDUser FROM User WHERE Phone = ? AND Pin = MD5(?)"
	err := db.QueryRow(query, &lgn.Phone, &lgn.Pin).Scan(&lgn.IDUser)
	return err
}
