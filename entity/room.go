package entity

import "time"

type Room struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	IsPrivate bool       `json:"is_private"`
	CreatedAt *time.Time `json:"created_at"`
}
