package handler

import (
	"encoding/json"
	"io"
	"log"

	"golang.org/x/net/websocket"

	"chatty/handler/bidirectional"
	"chatty/service"
	"chatty/util"
)

type websocketHandler struct {
	hub         *bidirectional.Hub
	chatService service.ChatService
}

func NewWebsocketHandler(
	hub *bidirectional.Hub,
	chatService service.ChatService,
) WebsocketHandler {
	return &websocketHandler{hub, chatService}
}

func (h *websocketHandler) Connect(conn *websocket.Conn) {
	// authenticate incoming connection
	username, ok := util.FromCtxUsername(conn.Request().Context())
	if !ok {
		log.Println("unauthenticated websocket connection")
		return
	}
	client := bidirectional.Client{
		Username: username,
		Conn:     conn,
	}
	h.hub.Connect(client)

	defer func(cl bidirectional.Client) {
		h.hub.Disconnect(cl)
	}(client)

	// listen to incoming messages
	for {
		var jsonBytes []byte
		if err := websocket.Message.Receive(conn, &jsonBytes); err != nil {
			if err != io.EOF {
				log.Println("encountered an error but not an EOF")
			}
			if err == websocket.ErrFrameTooLarge {
				log.Println("payload too large")
			}
			return
		}
		var message bidirectional.Message
		if err := json.Unmarshal(jsonBytes, &message); err != nil {
			log.Println("failed marshaling into Message")
			return
		}
		dbChat, err := h.chatService.Save(
			conn.Request().Context(),
			message.Chat,
		)
		if err != nil {
			log.Printf("failed storing chat: %s\n", err.Error())
			return
		}
		message.Chat = dbChat
		h.hub.Broadcast(message)
	}
}
