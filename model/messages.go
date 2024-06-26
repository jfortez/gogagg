package model

import "time"

type Message struct {
	Id         int       `json:"id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Status     string    `json:"status"`
	FromUserId int       `json:"fromUserId"`
	ToUserId   int       `json:"toUserId"`
}

type ChatMessage struct {
	MessageId        int       `json:"id"`
	MessageContent   string    `json:"content"`
	MessageCreatedAt time.Time `json:"createdAt"`
	MessageUpdatedAt time.Time `json:"updatedAt"`
	MessageStatus    string    `json:"status"`
	UserId           int       `json:"userId"`
	UserName         string    `json:"userName"`
	Avatar           string    `json:"userAvatar"`
}

type RequestedMessages struct {
	MessageId  int       `json:"messageId"`
	UserId     int       `json:"userId"`
	UserName   string    `json:"userName"`
	UserAvatar string    `json:"userAvatar"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type RequestMessage struct {
	Content             string    `json:"content"`
	UserId              int       `json:"userId"`
	UserName            string    `json:"userName"`
	LastInteractionTime time.Time `json:"lastInteractionTime"`
}

type CreateMessage struct {
	Content    string `json:"content"`
	Status     string `json:"status"`
	FromUserId int    `json:"fromUserId"`
	ToUserId   int    `json:"toUserId"`
}

type CurrentChatUser struct {
	UserId              int       `json:"userId"`
	UserName            string    `json:"userName"`
	LastInteractionTime time.Time `json:"lastInteractionTime"`
	Avatar              string    `json:"userAvatar"`
}

type AuthUser struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
