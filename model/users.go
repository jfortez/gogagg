package model

import "time"

// https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go
type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}
