package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/jfortez/gogagg/model"
	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	Connection *sql.DB
}

func seedUsers() []model.User {
	Users := []model.User{
		{Name: "John Doe", Email: "jdoe@doemail.com", Age: 23, Img: "https://i.pravatar.cc/150", Password: "123456", Description: "A simple todo app built with Go and Tailwind CSS"},
	}
	return Users
}

func New() *DataBase {
	os.Remove("./db/sql.db")

	db, err := sql.Open("sqlite3", "./db/sql.db")
	if err != nil {
		panic(err)
	}

	return &DataBase{Connection: db}
}

func (d *DataBase) InitDB() {
	sqlStmt := `
	DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		age INTEGER,
		img TEXT,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		password TEXT NOT NULL,
		description TEXT
	);
	DELETE FROM users;
	`
	_, err := d.Connection.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	tx, err := d.Connection.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into users(name,email,age,img, password, description) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	users := seedUsers()
	for _, v := range users {
		_, err = stmt.Exec(v.Name, v.Email, v.Age, v.Img, v.Password, v.Description)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DataBase) Close() {
	if d.Connection != nil {
		d.Connection.Close()
	}
}
