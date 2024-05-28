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
		{Id: 1, Name: "John Doe", Email: "jdoe@doemail.com", Age: 23, Img: "https://i.pravatar.cc/150", Password: "123456", Description: "A simple todo app built with Go and Tailwind CSS"},
		{Id: 2, Name: "Bonnie Green", Email: "bgreen@doemail.com", Age: 23, Img: "https://i.pravatar.cc/150", Password: "123456", Description: "A simple todo app built with Go and Tailwind CSS"},
		{Id: 3, Name: "John Smith", Email: "jsmith@doemail.com", Age: 23, Img: "https://i.pravatar.cc/150", Password: "123456", Description: "A simple todo app built with Go and Tailwind CSS"},
	}
	return Users
}
func seedMessages() []model.Message {
	Messages := []model.Message{
		{Content: "Hello", FromUserId: 1, ToUserId: 2, Status: "delivered"},
		{Content: "Hi", FromUserId: 2, ToUserId: 1, Status: "delivered"},
		{Content: "How are you?", FromUserId: 1, ToUserId: 2, Status: "delivered"},
		{Content: "I'm fine", FromUserId: 2, ToUserId: 1, Status: "delivered"},
		{Content: "What's up?", FromUserId: 1, ToUserId: 2, Status: "delivered"},
		{Content: "Nothing much", FromUserId: 2, ToUserId: 1, Status: "delivered"},
		{Content: "I'm good", FromUserId: 1, ToUserId: 2, Status: "delivered"},
		{Content: "Good to see you", FromUserId: 2, ToUserId: 1, Status: "delivered"},
		{Content: "Hey", FromUserId: 1, ToUserId: 3, Status: "pending"},
		{Content: "How are you?", FromUserId: 3, ToUserId: 1, Status: "pending"},
		{Content: "Hey", FromUserId: 3, ToUserId: 2, Status: "pending"},
	}
	return Messages
}

func New() *DataBase {
	os.Remove("./db/sql.db")

	db, err := sql.Open("sqlite3", "file:./db/sql.db?_foreign_keys=1&_pragma=foreign_keys(1)")
	// file:test.db?_foreign_keys=true
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

	DROP TABLE IF EXISTS messages;
	CREATE TABLE messages (
		id INTEGER NOT NULL PRIMARY KEY,
		content TEXT NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		status TEXT NOT NULL DEFAULT 'pending',
		fromUserId INTEGER NOT NULL,
		toUserId INTEGER NOT NULL
	);

	DELETE FROM users;
	DELETE FROM messages;
	`
	// ! add foreign keys later, there are some issues with sqlite and foreign keys
	// FOREIGN KEY (fromUserId) REFERENCES users(id),
	// FOREIGN KEY (toUserId) REFERENCES users(id)

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

	tx, err = d.Connection.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err = tx.Prepare("insert into messages(content, fromUserId, toUserId, status) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	messages := seedMessages()
	for _, v := range messages {
		_, err = stmt.Exec(v.Content, v.FromUserId, v.ToUserId, v.Status)
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
