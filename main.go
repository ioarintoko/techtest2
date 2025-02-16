package main

import (
	"fmt"
	"net/http"
	"techtest2/handler"
	"techtest2/lib"
)

var database string

func init() {
	database = "techtest2"
}

func main() {
	db, err := lib.ConnectMySql(database)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	handler.RegisDB(db)
	http.HandleFunc("/api/", handler.API)
	http.ListenAndServe(":8087", nil)
}
