package test

import (
	"techtest2/lib"
	"techtest2/model"
	"testing"
)

var database, databaseDefaultMysql string

func init() {
	database = "techtest2"
	databaseDefaultMysql = "Mysql"
}

func TestDatabaseMysql(t *testing.T) {
	t.Run("MySQL Connection Testing", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()
	})

	t.Run("Drop Table Testing", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.DropDB(db, database)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create DB Testing", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateDB(db, database)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table User", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TableUser)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table Transaction", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TableTransaction)
		if err != nil {
			t.Fatal(err)
		}
	})
}
