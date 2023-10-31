package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func Setup() {
	var err error
	DB, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic("Could not connect to db")
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Could not ping db", err)
	}
}
