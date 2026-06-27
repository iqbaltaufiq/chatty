package entity

import "time"

type Chat struct {
	ID        int        `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Message   string     `json:"message,omitempty"`
	RoomID    string     `json:"room_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
