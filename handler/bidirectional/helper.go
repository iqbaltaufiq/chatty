package bidirectional

import (
	"log"

	"golang.org/x/net/websocket"
)

func sendMessage(h *hub, message Message) {
	// broadcast this message to the audience
	for _, aud := range message.Audiences {
		conn, exists := h.clients[aud]
		if !exists {
			continue
		}
		if err := websocket.Message.Send(
			conn,
			message.Stringify(),
		); err != nil {
			log.Printf("broadcast error: %s\n", err.Error())
			continue
		}
	}
}
