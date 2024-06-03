package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "naol:naol.@tcp(localhost:3306)/ForTest")
	if err != nil {
		fmt.Println("error validating database")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error validating database")
		panic(err.Error())
	}
	insert, err := db.Query("INSERT INTO ForTest.userTest VALUES('kasu', 45)")
	if err != nil {
		fmt.Println("error inserting data")
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully inserted data")
}