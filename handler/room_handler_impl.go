package handler

import (
	"encoding/json"
	"net/http"

	"chatty/entity"
	"chatty/service"
	"chatty/util"
)

type roomHandler struct {
	roomService service.RoomService
}

func NewRoomHandler(roomService service.RoomService) RoomHandler {
	return &roomHandler{roomService}
}

func (h *roomHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}

	var room entity.Room
	room.Name = payload.Name
	if _, err := h.roomService.Create(r.Context(), room); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusCreated, util.M{
		"message": "A room has been created.",
	})
}

func (h *roomHandler) UpdateName(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		RoomID string `json:"room_id"`
		Name   string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if err := h.roomService.UpdateName(
		r.Context(),
		payload.RoomID,
		payload.Name,
	); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"message": "Room has been updated successfully.",
	})
}
