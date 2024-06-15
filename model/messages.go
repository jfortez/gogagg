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
	FromUserId       int       `json:"fromUserId"`
	FromUserName     string    `json:"fromUserName"`
	ToUserId         int       `json:"toUserId"`
	ToUserName       string    `json:"toUserName"`
}

type RequestedMessages struct {
	MessageId int       `json:"messageId"`
	UserId    int       `json:"userId"`
	UserName  string    `json:"userName"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RequestMessage struct {
	Content string `json:"content"`
}

type CreateMessage struct {
	Content    string `json:"content"`
	Status     string `json:"status"`
	FromUserId int    `json:"fromUserId"`
	ToUserId   int    `json:"toUserId"`
}
