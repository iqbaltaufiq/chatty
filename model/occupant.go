package model

import "time"

type Occupant struct {
	ID        int        `json:"id"`
	RoomID    string     `json:"room_id"`
	Username  string     `json:"username"`
	CreatedAt *time.Time `json:"created_at"`
}
