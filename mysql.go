package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_dev")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
