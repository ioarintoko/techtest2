package model

import (
	"database/sql"
	"fmt"
	"strings"
	"techtest2/lib"
	"time"

	"github.com/google/uuid"
)

type TransactionResponse struct {
	IDAlias      string    `json:"idalias"`
	Amount       int       `json:"amount"`
	Remarks      string    `json:"remarks"`
	Status       string    `json:"status"`
	BalanceStart int       `json:"balancestart"`
	BalanceEnd   int       `json:"balanceend"`
	CreateDate   time.Time `json:"createdate"`
}

type Transaction struct {
	IDTransaction   string    `json:"idtransaction"`
	IDAlias         string    `json:"idalias"`
	IDUser          string    `json:"iduser"`
	IDReference     string    `json:"idreference"`
	TransactionType string    `json:"transactiontype"`
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	Amount          int       `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceStart    int       `json:"balancestart"`
	BalanceEnd      int       `json:"balanceend"`
	CreateDate      time.Time `json:"createdate"`
	UpdateDate      time.Time `json:"updatedate"`
}

var TableTransaction = lib.Table{
	Name: "Transaction",
	Field: []string{
		"IDTransaction VARCHAR(30) PRIMARY KEY",
		"IDAlias VARCHAR(30)",
		"IDUser VARCHAR(30)",
		"IDReference VARCHAR(30)",
		"Type VARCHAR(50)",
		"TransactionType VARCHAR(50)",
		"Amount INT(20)",
		"Remarks TEXT",
		"Status VARCHAR(20)",
		"BalanceStart INT(20)",
		"BalanceEnd INT(20)",
		"CreateDate DATETIME",
		"UpdateDate DATETIME",
	},
}

func (t *Transaction) Insert(db *sql.DB) (*TransactionResponse, error) {
	newUUID := uuid.New()
	now := time.Now()
	alias := ""
	if t.Type == "Payment" {
		alias = "payment_id"
	} else if t.Type == "topup" {
		alias = "top_up_id"
	} else {
		alias = "transfer_id"
	}

	query := `INSERT INTO Transaction (IDTransaction, IDAlias, IDUser, IDReference, TransactionType, Type, Amount, Remarks, Status, BalanceStart, BalanceEnd, CreateDate, UpdateDate)
				VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)`

	_, err := db.Exec(query, newUUID.String(), alias, t.IDUser, t.IDReference, t.TransactionType, t.Type, t.Amount, t.Remarks, t.Status, t.BalanceStart, t.BalanceEnd, now, now)
	if err != nil {
		return nil, err
	}

	// Return inserted user as a struct
	insertedTransaction := &TransactionResponse{
		IDAlias:      t.IDAlias,
		Amount:       t.Amount,
		Remarks:      t.Remarks,
		Status:       t.Status,
		BalanceStart: t.BalanceStart,
		BalanceEnd:   t.BalanceEnd,
		CreateDate:   t.CreateDate,
	}

	return insertedTransaction, nil
}

func (t *Transaction) Update(db *sql.DB, datatr map[string]interface{}) (*TransactionResponse, error) {
	var kolom []string
	var args []interface{}

	// Buat daftar kolom untuk UPDATE
	for key, value := range datatr {
		if value == "" {
			continue
		}
		kolom = append(kolom, fmt.Sprintf("%s = ?", strings.ToLower(key)))
		args = append(args, value)
	}

	// Jika tidak ada kolom yang diperbarui, hentikan
	if len(kolom) == 0 {
		return nil, fmt.Errorf("tidak ada data yang diperbarui")
	}

	// Buat query update tanpa backtick di SET
	dataUpdate := strings.Join(kolom, ", ")
	query := fmt.Sprintf("UPDATE `Transaction` SET %s WHERE IDTransaction = ?", dataUpdate)
	args = append(args, t.IDTransaction) // Tambahkan IDTransaction sebagai argumen terakhir

	// Debug query
	fmt.Println("Executing query:", query, args)

	// Eksekusi query update
	_, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	// Ambil data terbaru dari database setelah update
	var updated TransactionResponse
	err = db.QueryRow("SELECT IDAlias, Amount, Remarks, Status, BalanceStart, BalanceEnd, CreateDate FROM `Transaction` WHERE IDTransaction = ?", t.IDTransaction).
		Scan(&updated.IDAlias, &updated.Amount, &updated.Remarks, &updated.Status, &updated.BalanceStart, &updated.BalanceEnd, &updated.CreateDate)

	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func GetsTransaction(db *sql.DB, params ...string) ([]*Transaction, error) {
	var kolom []string
	var args []interface{}

	// Cek jika ada parameter sebelum mengakses
	if len(params) > 0 && params[0] != "" {
		dataParams := strings.Split(params[len(params)-1], ";")

		for _, val := range dataParams {
			temp := strings.Split(val, ",")
			if len(temp) < 3 { // Cek apakah temp memiliki minimal 3 elemen
				continue // Skip jika format tidak valid
			}

			column := strings.ToLower(temp[0]) // Pastikan nama kolom dalam huruf kecil
			operator := temp[1]
			value := temp[2]

			where := fmt.Sprintf("%s %s ?", column, operator)
			kolom = append(kolom, where)
			args = append(args, value)
		}
	}

	// Bangun query
	query := "SELECT * FROM `Transaction`"
	if len(kolom) > 0 {
		query += " WHERE " + strings.Join(kolom, " AND ")
	}

	// Debug query
	fmt.Println("Executing query:", query, args)

	// Eksekusi query
	datatr, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer datatr.Close()

	// Parsing hasil query
	var result []*Transaction
	for datatr.Next() {
		each := &Transaction{}
		err := datatr.Scan(
			&each.IDTransaction, &each.IDAlias, &each.IDUser, &each.IDReference,
			&each.TransactionType, &each.Type, &each.Amount, &each.Remarks,
			&each.Status, &each.BalanceStart, &each.BalanceEnd, &each.CreateDate, &each.UpdateDate,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	// Tambahkan log jika tidak ada hasil
	if len(result) == 0 {
		fmt.Println("No transactions found")
	}

	return result, nil
}
