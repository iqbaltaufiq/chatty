package handler

import (
	"net/http"

	"chatty/service"
)

type chatHandler struct {
	chatService service.ChatService
}

func NewChatHandler(chatService service.ChatService) ChatHandler {
	return &chatHandler{chatService}
}

func (h *chatHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// wip
}

func (h *chatHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// wip
}
