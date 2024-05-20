package services

import (
	"database/sql"
	"log"

	"github.com/jfortez/gogagg/model"
)

func GetUsers(connection *sql.DB) (users []model.User) {
	// users = make([]User, 0)
	rows, err := connection.Query("SELECT id,name,email,age,img,createdAt FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Img, &user.CreatedAt)
		users = append(users, user)
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

func GetUser(connection *sql.DB, id string) (user model.User) {

	stmt, err := connection.Prepare("SELECT id,name,email,age, img,createdAt FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Img, &user.CreatedAt)
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
	stmt, err := connection.Prepare("UPDATE users SET name=?, email=?, age=? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Age, user.Id)
	if err != nil {
		log.Fatal(err)
	}
}
