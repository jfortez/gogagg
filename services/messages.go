package services

import (
	"database/sql"

	"github.com/jfortez/gogagg/model"
)

func SendMessage(connection *sql.DB, message model.Message) error {
	stmt, err := connection.Prepare("INSERT INTO messages(content,createdAt,updatedAt,status,fromUserId,toUserId) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(message.Content, message.CreatedAt, message.UpdatedAt, message.Status, message.FromUserId, message.ToUserId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMessageStatus(connection *sql.DB, message model.Message) error {
	stmt, err := connection.Prepare("UPDATE messages SET status=? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(message.Status, message.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetMessages(connection *sql.DB, fromUserId int, toUserId int) (messages []model.Message) {
	rows, err := connection.Query("SELECT id,content,createdAt,updatedAt,status,fromUserId,toUserId FROM messages WHERE fromUserId = ? AND toUserId = ?", fromUserId, toUserId)
	if err != nil {
		return messages
	}
	defer rows.Close()
	for rows.Next() {
		var message model.Message
		err = rows.Scan(&message.Id, &message.Content, &message.CreatedAt, &message.UpdatedAt, &message.Status, &message.FromUserId, &message.ToUserId)
		messages = append(messages, message)
		if err != nil {
			return messages
		}
	}

	return
}

func GetMessageListByCurrentUser(connection *sql.DB, userId int) (messages []model.Message) {
	rows, err := connection.Query("SELECT id,content,createdAt,updatedAt,status,fromUserId,toUserId FROM messages WHERE toUserId = ?", userId)
	if err != nil {
		return messages
	}
	defer rows.Close()
	for rows.Next() {
		var message model.Message
		err = rows.Scan(&message.Id, &message.Content, &message.CreatedAt, &message.UpdatedAt, &message.Status, &message.FromUserId, &message.ToUserId)
		messages = append(messages, message)
		if err != nil {
			return messages
		}
	}

	return
}
