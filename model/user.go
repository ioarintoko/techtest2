package model

import (
	"database/sql"
	"fmt"
	"strings"
	"techtest2/lib"
	"time"

	"github.com/google/uuid"
)

type UserRegisterResponse struct {
	IDUser     string    `json:"iduser"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	CreateDate time.Time `json:"createdate"`
}

type UserUpdateResponse struct {
	IDUser     string    `json:"iduser"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Address    string    `json:"address"`
	ModifyDate time.Time `json:"modifydate"`
}

type User struct {
	IDUser     string    `json:"iduser"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	Pin        string    `json:"pin"`
	CreateDate time.Time `json:"createdate"`
	ModifyDate time.Time `json:"modifydate"`
}

var TableUser = lib.Table{
	Name: "User",
	Field: []string{
		"IDUser VARCHAR(30) PRIMARY KEY",
		"FirstName VARCHAR(50)",
		"LastName VARCHAR(50)",
		"Phone VARCHAR(20)",
		"Address TEXT",
		"Pin TEXT",
		"CreateDate DATETIME",
		"ModifyDate DATETIME",
	},
}

func (u *User) Insert(db *sql.DB) (*UserRegisterResponse, error) {
	newUUID := uuid.New()
	now := time.Now()
	query := `INSERT INTO User (IDUser, FirstName, LastName, Phone, Address, Pin, CreateDate, ModifyDate)
				VALUES(?,?,?,?,?,MD5(?),?,?)`

	_, err := db.Exec(query, newUUID.String(), u.FirstName, u.LastName, u.Phone, u.Address, u.Pin, now, now)
	if err != nil {
		return nil, err
	}

	// Return inserted user as a struct
	insertedUser := &UserRegisterResponse{
		IDUser:     newUUID.String(),
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Phone:      u.Phone,
		Address:    u.Address,
		CreateDate: now,
	}

	return insertedUser, nil
}

func (u *User) Delete(db *sql.DB) error {
	query := "DELETE FROM User WHERE IDUser = ?"
	_, err := db.Exec(query, u.IDUser)
	return err
}

func (u *User) Update(db *sql.DB, datauser map[string]interface{}) (*UserUpdateResponse, error) {
	var kolom = []string{}
	var args []interface{}
	now := time.Now()

	for key, value := range datauser {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ", ")
	query := fmt.Sprintf("UPDATE User SET %s WHERE IDUser = ?", dataUpdate) // Use ? as a placeholder

	args = append(args, u.IDUser)     // Append IDUser at the end
	_, err := db.Exec(query, args...) // Pass args safely

	if err != nil {
		return nil, err
	}

	// Return updated user data
	updatedUser := &UserUpdateResponse{
		IDUser:     u.IDUser,
		FirstName:  u.FirstName, // Gunakan nilai lama jika tidak diperbarui
		LastName:   u.LastName,
		Address:    u.Address,
		ModifyDate: now,
	}

	if val, ok := datauser["FirstName"]; ok {
		updatedUser.FirstName = val.(string)
	}

	if val, ok := datauser["LastName"]; ok {
		updatedUser.LastName = val.(string)
	}

	if val, ok := datauser["Address"]; ok {
		updatedUser.Address = val.(string)
	}

	return updatedUser, nil
}

func (u *User) Get(db *sql.DB) error {
	query := "SELECT * FROM User WHERE IDUser = ?"
	err := db.QueryRow(query, u.IDUser).Scan(&u.IDUser, &u.FirstName, &u.LastName, &u.Phone, &u.Address,
		&u.Pin, &u.CreateDate, &u.ModifyDate)
	return err
}

func (u *User) Profile(db *sql.DB) (*User, error) {
	query := "SELECT * FROM User WHERE IDUser = ?"
	datauser, err := db.Query(query, &u.IDUser)
	if err != nil {
		return nil, err
	}
	defer datauser.Close()

	var result *User
	for datauser.Next() {
		each := &User{}
		err = datauser.Scan(&each.IDUser, &each.FirstName, &each.LastName, &each.Phone, &each.Address,
			&each.Pin, &each.CreateDate, &each.ModifyDate)
		if err != nil {
			return nil, err
		}
		result = each
	}

	return result, nil
}

func GetsUser(db *sql.DB, params ...string) ([]*User, error) {
	var kolom = []string{}
	var args []interface{}

	if len(params) != 0 {
		if params[0] != "" {
			dataParams := strings.Split(params[len(params)-1], ";")
			for _, val := range dataParams {
				temp := strings.Split(fmt.Sprintf("%s", val), ",")
				where := fmt.Sprintf("%s %s ?", strings.ToLower(temp[0]), temp[1])
				kolom = append(kolom, where)
				args = append(args, temp[2])
			}
		}
	}

	dataKondisi := strings.Join(kolom, " AND ")
	var query string
	query = "SELECT * FROM User"
	if dataKondisi != "" {
		query = "SELECT * FROM User WHERE " + dataKondisi
	}

	datauser, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer datauser.Close()

	var result []*User
	for datauser.Next() {
		each := &User{}
		err := datauser.Scan(&each.IDUser, &each.FirstName, &each.LastName, &each.Phone, &each.Address,
			&each.Pin, &each.CreateDate, &each.ModifyDate)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}
