package handler

import (
	"net/http"

	"chatty/service"
	"chatty/util"
)

type chatHandler struct {
	chatService service.ChatService
}

func NewChatHandler(chatService service.ChatService) ChatHandler {
	return &chatHandler{chatService}
}

func (h *chatHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	roomID := r.PathValue("roomID")

	chats, err := h.chatService.GetAll(r.Context(), roomID)
	if err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"message":   "found",
		"rowsCount": len(chats),
		"chats":     chats,
	})
}

func (h *chatHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// wip
}
