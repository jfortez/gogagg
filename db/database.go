package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Storage struct {
	*sql.DB
}

func New() *Storage {

	PG_USER := os.Getenv("DB_USER")
	PG_PASSWORD := os.Getenv("DB_PASSWORD")
	PG_DATABASE := os.Getenv("DB_NAME")
	PG_HOST := "localhost"
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", PG_USER, PG_PASSWORD, PG_HOST, PG_DATABASE)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return &Storage{
		DB: db,
	}
}
