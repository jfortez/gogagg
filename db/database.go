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
		{Name: "John Doe", Email: "jdoe@doemail.com", Age: 23, Img: "https://i.pravatar.cc/150"},
		{Name: "Jane Doe", Email: "jadoe@doemail.com", Age: 23, Img: "https://i.pravatar.cc/120"},
	}
	return Users
}

func New() *DataBase {
	os.Remove("./sql.db")

	db, err := sql.Open("sqlite3", "./sql.db")
	if err != nil {
		panic(err)
	}

	database := &DataBase{Connection: db}
	database.initDB()
	return database
}

func (d *DataBase) initDB() {
	sqlStmt := `
	CREATE TABLE users (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER,
		img TEXT,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
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
	stmt, err := tx.Prepare("insert into users(name,email,age, img) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	users := seedUsers()
	for _, v := range users {
		_, err = stmt.Exec(v.Name, v.Email, v.Age, v.Img)
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
