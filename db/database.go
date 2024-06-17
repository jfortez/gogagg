package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	*sql.DB
}

func New() *Storage {

	connStr := "postgres://root:root@localhost/gogag?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return &Storage{
		DB: db,
	}
}
