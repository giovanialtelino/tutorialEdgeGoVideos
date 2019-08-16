package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Tester struct {
	ID       int    `json:"id"`
	name     string `json:"name"`
	whatever string `json:"whatever"`
}

func main() {
	fmt.Println("Go Postgre adapted")

	connStr := "user=unis password=patoverde dbname=whatever"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO tester VALUES(2, 'GIOVANI2', 'OFF2')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Inserted table")

	results, err := db.Query("SELECT Id, name, whatever FROM tester")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tester Tester
		err = results.Scan(&tester.ID, &tester.name, &tester.whatever)
		if err != nil {
			panic(err.Error())
		}

		fmt.Print(tester.name)
		fmt.Print(tester.whatever)
	}

}
