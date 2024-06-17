package services

import (
	"database/sql"

	"github.com/jfortez/gogagg/model"
)

func SendMessage(connection *sql.DB, message model.CreateMessage) error {
	stmt, err := connection.Prepare("INSERT INTO messages(content,status,fromUserId,toUserId) VALUES($1,$2,$3,$4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(message.Content, message.Status, message.FromUserId, message.ToUserId)
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

func GetMessages(connection *sql.DB, fromUserId int, toUserId int) (messages []model.ChatMessage, err error) {

	query := `
		SELECT 
				m.id AS messageId,
				m.content AS messageContent,
				m.createdAt AS messageCreatedAt,
				m.updatedAt AS messageUpdatedAt,
				m.status AS messageStatus,
				m.fromUserId,
				u_from.name AS fromUserName,
				u_from.avatar AS userAvatar
		FROM 
				messages m
		JOIN 
				users u_from ON m.fromUserId = u_from.id
		JOIN 
				users u_to ON m.toUserId = u_to.id
		WHERE 
				(m.fromUserId = $1 AND m.toUserId = $2) 
				OR 
				(m.fromUserId = $2 AND m.toUserId = $1)
		ORDER BY 
				m.createdAt ASC;
`
	rows, err := connection.Query(query, fromUserId, toUserId)
	if err != nil {
		return messages, err
	}
	defer rows.Close()
	for rows.Next() {
		var message model.ChatMessage
		err = rows.Scan(&message.MessageId, &message.MessageContent, &message.MessageCreatedAt, &message.MessageUpdatedAt, &message.MessageStatus, &message.UserId, &message.UserName, &message.Avatar)
		messages = append(messages, message)
		if err != nil {
			return messages, err
		}
	}

	return
}

func GetMessageListByCurrentUser(connection *sql.DB, userId int) (messages []model.RequestedMessages, err error) {
	query := `
SELECT 
		m.id as messageId,
    u.id AS userId,
    u.name AS userName,
		u.avatar AS userAvatar,
    m.content AS lastMessageContent,
    m.createdAt AS lastMessageCreatedAt,
    m.updatedAt AS lastMessageUpdatedAt
FROM 
    users u
JOIN 
    (SELECT 
         DISTINCT ON (fromUserId) 
				 id,
         fromUserId, 
         content, 
         createdAt, 
         updatedAt 
     FROM 
         messages 
     WHERE 
         toUserId = $1
     ORDER BY 
         fromUserId, 
         id DESC 
    ) m 
ON 
    u.id = m.fromUserId
ORDER BY 
    m.createdAt DESC;
`

	rows, err := connection.Query(query, userId)

	if err != nil {
		return messages, err
	}

	defer rows.Close()

	for rows.Next() {
		var message model.RequestedMessages
		err = rows.Scan(&message.MessageId, &message.UserId, &message.UserName, &message.UserAvatar, &message.Content, &message.CreatedAt, &message.UpdatedAt)
		messages = append(messages, message)
		if err != nil {
			return messages, err
		}
	}

	return
}
