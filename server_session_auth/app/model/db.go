package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Setup() {

	dsn := "host=172.17.0.2 port=5432 user=admin password=test dbname=admin sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}