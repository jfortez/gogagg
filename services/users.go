package services

import (
	"database/sql"
	"log"

	"github.com/jfortez/gogagg/model"
)

func CreateUser(connection *sql.DB, user model.User) error {
	stmt, err := connection.Prepare("INSERT INTO users(name,email,age,img, password, description) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Age, user.Avatar, user.Password, user.Description)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(connection *sql.DB) (users []model.User) {
	rows, err := connection.Query("SELECT id,name,email,age,img,password,description,createdAt FROM users WHERE password IS NOT NULL")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Avatar, &user.Password, &user.Description, &user.CreatedAt)
		users = append(users, user)
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

func FindAuthUser(connection *sql.DB, email string, password string) (user model.User, err error) {
	stmt, err := connection.Prepare("SELECT id,name,email,age,img,description,createdAt FROM users WHERE email = ? AND password = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email, password)

	if row.Err() != nil {
		return user, row.Err()
	}

	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Avatar, &user.Description, &user.CreatedAt)

	if err != nil {
		return user, err
	}

	return
}

func GetUser(connection *sql.DB, id string) (user model.User) {

	stmt, err := connection.Prepare("SELECT id,name,email,age,img,password,description,createdAt FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Avatar, &user.Password, &user.Description, &user.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func DeleteUser(connection *sql.DB, id string) {
	stmt, err := connection.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateUser(connection *sql.DB, user model.User) {
	stmt, err := connection.Prepare("UPDATE users SET name=?, email=?, age=?, password=?, description=? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Age, user.Id, user.Password, user.Description)
	if err != nil {
		log.Fatal(err)
	}
}
