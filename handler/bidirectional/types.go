package bidirectional

import (
	"encoding/json"
	"log"
	"time"
)

type Message struct {
	ID        int        `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	RoomID    string     `json:"room_id,omitempty"`
	Message   string     `json:"message,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (m *Message) Stringify() string {
	strMessage, err := json.Marshal(m)
	if err != nil {
		log.Printf("Marshal error: %s\n", err.Error())
		return ""
	}
	return string(strMessage)
}
