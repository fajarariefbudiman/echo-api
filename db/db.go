package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewData() *sql.DB {
	db, err := sql.Open("mysql", "fajar:fajararief2006@tcp(localhost:3306)/echo_rest")
	if err != nil {
		log.Println("Connection to Database Error")
		panic(err)
	}

	return db
}
