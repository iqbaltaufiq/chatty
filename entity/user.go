package entity

import "time"

type User struct {
	UUID      string     `json:"uuid,omitempty"`
	Username  string     `json:"username,omitempty"`
	Fullname  string     `json:"fullname,omitempty"`
	Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	Photo     string     `json:"photo,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
