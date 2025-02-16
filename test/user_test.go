package test

import (
	"database/sql"
	"techtest2/lib"
	"techtest2/model"
	"testing"
)

var datauser = []*model.User{
	{
		FirstName: "Bramantio",
		LastName:  "Galih",
		Phone:     "081998998155",
		Address:   "Jakarta",
		Pin:       "Bram9090.",
	},
	{
		FirstName: "Norman",
		LastName:  "Osbourne",
		Phone:     "081998998155",
		Address:   "Jakarta",
		Pin:       "Bram9090.",
	},
}

func ConnectDB(t *testing.T) (*sql.DB, error) {
	db, err := lib.ConnectMySql(database)
	if err != nil {
		t.Fatal(err)
		return nil, err
	}

	return db, nil
}

func TestUser(t *testing.T) {
	t.Run("Test Insert Table User", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		for _, val := range datauser {
			_, err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table User", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataUpdate := map[string]interface{}{
			"LastName": "Osbourne",
		}

		_, err := datauser[0].Update(db, dataUpdate)
		if err != nil {
			t.Fatal(err)
		}
	})
}
