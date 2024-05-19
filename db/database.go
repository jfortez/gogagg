package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/jfortez/gogagg/model"

	_ "github.com/mattn/go-sqlite3"
)

var Pool *sql.DB

func seedUsers() []model.User {
	Users := []model.User{
		{Name: "John Doe", Email: "jdoe@doemail.com", Age: 23},
		{Name: "Jane Doe", Email: "jadoe@doemail.com", Age: 23},
	}
	return Users
}

func InitDB() {

	os.Remove("./sql.db")

	var err error
	Pool, err = sql.Open("sqlite3", "./sql.db")
	if err != nil {
		panic(err)
	}
	// defer Pool.Close()

	sqlStmt := `
	CREATE TABLE users (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
	);
	DELETE FROM users;
	`
	_, err = Pool.Exec(sqlStmt)

	if err != nil {
		panic(err)
	}

	tx, err := Pool.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into users(name,email,age) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	users := seedUsers()

	for _, v := range users {
		_, err = stmt.Exec(v.Name, v.Email, v.Age)
		if err != nil {
			log.Fatal(err)
		}

	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
