package handler

import "net/http"

type RoomHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	UpdateName(w http.ResponseWriter, r *http.Request)
}
