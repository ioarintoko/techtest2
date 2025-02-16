package lib

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Table struct {
	Name  string
	Field []string
}

type ForeignKey struct {
	Name         string
	ForeignName  string
	Field        string
	ForeignField string
}

func ConnectMySql(database string) (*sql.DB, error) {
	connStr := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?parseTime=true", database)
	db, err := sql.Open("mysql", connStr)
	return db, err
}

func DropDB(db *sql.DB, name string) error {
	query := fmt.Sprintf("DROP DATABASE %s", name)
	_, err := db.Exec(query)
	return err
}

func CreateDB(db *sql.DB, name string) error {
	query := fmt.Sprintf("CREATE DATABASE %s", name)
	_, err := db.Exec(query)
	return err
}

func CreateTable(db *sql.DB, table Table) error {
	query := fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, strings.Join(table.Field, ","))
	_, err := db.Exec(query)
	return err
}

func AddForeignKey(db *sql.DB, fk ForeignKey) error {
	query := fmt.Sprintf(
		"ALTER TABLE %s ADD FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE RESTRICT ON UPDATE CASCADE",
		fk.Name, fk.Field, fk.ForeignName, fk.ForeignField)
	_, err := db.Exec(query)
	return err
}
