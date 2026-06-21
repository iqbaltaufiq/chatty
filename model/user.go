package model

import "time"

type User struct {
	UUID      string     `json:"uuid"`
	Username  string     `json:"username"`
	Fullname  string     `json:"fullname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Photo     string     `json:"photo"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
