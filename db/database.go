package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataBase struct {
	Connection *sql.DB
}

func New() *DataBase {

	connStr := "postgres://root:root@localhost/gogag?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return &DataBase{Connection: db}
}
