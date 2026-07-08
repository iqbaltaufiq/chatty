package handler

import (
	"net/http"

	"chatty/service"
)

type roomHandler struct {
	roomService service.RoomService
}

func NewRoomHandler(roomService service.RoomService) RoomHandler {
	return &roomHandler{roomService}
}

func (h *roomHandler) Create(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}

func (h *roomHandler) UpdateName(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}
