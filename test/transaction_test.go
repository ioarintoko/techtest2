package test

import (
	"fmt"
	"techtest2/model"
	"testing"
)

var datatr = []*model.Transaction{
	{
		IDAlias:         "top_up_id",
		IDUser:          "",
		IDReference:     "",
		TransactionType: "KREDIT",
		Type:            "Topup",
		Status:          "Pending",
		Amount:          50000,
		Remarks:         "",
		BalanceStart:    0,
		BalanceEnd:      50000,
	},
	{
		IDAlias:         "payment_id",
		IDUser:          "",
		IDReference:     "",
		TransactionType: "DEBIT",
		Type:            "Payment",
		Status:          "Pending",
		Amount:          10000,
		Remarks:         "Bayar Listrik",
		BalanceStart:    50000,
		BalanceEnd:      40000,
	},
}

func TestTransaction(t *testing.T) {
	t.Run("Test Insert Table Transaction", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		for _, val := range datatr {
			_, err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table Transaction", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		dataUpdate := map[string]interface{}{
			"status": "Success",
		}

		_, err := datauser[0].Update(db, dataUpdate)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets Table Transaction", func(t *testing.T) {
		db, _ := ConnectDB(t)
		defer db.Close()

		datatr, err := model.GetsTransaction(db)
		if err != nil {
			t.Fatal(err)
		}

		for _, val := range datatr {
			fmt.Println(*val)
		}
	})
}
