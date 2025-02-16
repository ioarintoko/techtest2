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
	Amount       string    `json:"amount"`
	Remarks      string    `json:"remarks"`
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
	Amount          string    `json:"amount"`
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
		"IDAlias VARCHAR(30) PRIMARY KEY",
		"IDUser VARCHAR(30) PRIMARY KEY",
		"IDReference VARCHAR(30) PRIMARY KEY",
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

	query := `INSERT INTO User (IDTransaction, IDAlias, IDUser, IDReference, Type, Amount, Remarks, Status, BalanceStart, BalanceEnd, CreateDate, UpdateDate)
				VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)`

	_, err := db.Exec(query, newUUID.String(), alias, t.IDUser, t.IDReference, t, t.TransactionType, t.Type, t.Amount, t.Remarks, t.BalanceStart, t.BalanceEnd, now, now)
	if err != nil {
		return nil, err
	}

	// Return inserted user as a struct
	insertedTransaction := &TransactionResponse{
		IDAlias:      t.IDAlias,
		Amount:       t.Amount,
		Remarks:      t.Remarks,
		BalanceStart: t.BalanceStart,
		BalanceEnd:   t.BalanceEnd,
		CreateDate:   t.CreateDate,
	}

	return insertedTransaction, nil
}

func GetsTransaction(db *sql.DB, params ...string) ([]*Transaction, error) {
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
	query = "SELECT * FROM Transaction"
	if dataKondisi != "" {
		query = "SELECT * FROM Transaction WHERE " + dataKondisi
	}

	datatr, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer datatr.Close()

	var result []*Transaction
	for datatr.Next() {
		each := &Transaction{}
		err := datatr.Scan(&each.IDTransaction, &each.IDAlias, &each.IDUser, &each.IDReference, &each.TransactionType, &each.Type, &each.Amount, &each.Remarks, &each.BalanceStart, &each.BalanceEnd, &each.CreateDate, &each.UpdateDate)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}
