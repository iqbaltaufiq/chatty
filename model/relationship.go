package model

import "time"

type Relationship struct {
	ID        int        `json:"id"`
	Initiator string     `json:"initiator"`
	Target    string     `json:"target"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
