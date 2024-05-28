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
