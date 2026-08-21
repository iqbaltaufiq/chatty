package handler

import (
	"encoding/json"
	"net/http"

	"chatty/service"
	"chatty/util"
)

type occupantHandler struct {
	occupantService service.OccupantService
}

func NewOccupantHandler(occupantService service.OccupantService) OccupantHandler {
	return &occupantHandler{occupantService}
}

func (h *occupantHandler) Add(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		RoomID   string `json:"room_id"`
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if payload.RoomID == "" || payload.Username == "" {
		util.JSONError(w, http.StatusBadRequest, &RequestErrInsufficientBody{})
		return
	}
	if err := h.occupantService.Add(
		r.Context(),
		payload.RoomID,
		payload.Username,
	); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.SendJSON(w, http.StatusOK, messagePayload{"A user has been added to the room."})
}

func (h *occupantHandler) Remove(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		RoomID   string `json:"room_id"`
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if payload.RoomID == "" || payload.Username == "" {
		util.JSONError(w, http.StatusBadRequest, &RequestErrInsufficientBody{})
		return
	}
	if err := h.occupantService.Remove(
		r.Context(),
		payload.RoomID,
		payload.Username,
	); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.SendJSON(w, http.StatusOK, messagePayload{"A user has been removed from the room."})
}
