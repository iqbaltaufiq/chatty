package bidirectional

import (
	"encoding/json"
	"log"

	"golang.org/x/net/websocket"

	"chatty/entity"
)

type Client struct {
	Conn     *websocket.Conn
	Username string
}

type Message struct {
	entity.Chat
	Audiences []string `json:"audiences,omitempty"`
}

func (m *Message) Stringify() string {
	strMessage, err := json.Marshal(m)
	if err != nil {
		log.Printf("Marshal error: %s\n", err.Error())
		return ""
	}
	return string(strMessage)
}
