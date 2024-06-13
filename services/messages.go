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

func GetMessages(connection *sql.DB, fromUserId int, toUserId int) (messages []model.Message, err error) {
	rows, err := connection.Query("SELECT id,content,createdAt,updatedAt,status,fromUserId,toUserId FROM messages WHERE fromUserId = ? AND toUserId = ?", fromUserId, toUserId)
	if err != nil {
		return messages, err
	}
	defer rows.Close()
	for rows.Next() {
		var message model.Message
		err = rows.Scan(&message.Id, &message.Content, &message.CreatedAt, &message.UpdatedAt, &message.Status, &message.FromUserId, &message.ToUserId)
		messages = append(messages, message)
		if err != nil {
			return messages, err
		}
	}

	return
}

func GetMessageListByCurrentUser(connection *sql.DB, userId int) (messages []model.RequestedMessages, err error) {

	rows, err := connection.Query("SELECT m.id, m.content, u.id, u.name FROM messages m JOIN users u ON m.fromuserid = u.id WHERE m.touserid = $1", userId)

	if err != nil {
		return messages, err
	}

	defer rows.Close()

	for rows.Next() {
		var message model.RequestedMessages
		err = rows.Scan(&message.MessageId, &message.Content, &message.RequestedByUserId, &message.RequestedUserName)
		messages = append(messages, message)
		if err != nil {
			return messages, err
		}
	}

	return
}
