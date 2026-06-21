package model

import "time"

type Chat struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Message   string     `json:"message"`
	RoomID    string     `json:"room_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
