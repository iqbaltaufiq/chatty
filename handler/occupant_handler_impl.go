package handler

import (
	"net/http"

	"chatty/service"
)

type occupantHandler struct {
	occupantService service.OccupantService
}

func NewOccupantHandler(occupantService service.OccupantService) OccupantHandler {
	return &occupantHandler{occupantService}
}

func (h *occupantHandler) Add(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}

func (h *occupantHandler) Remove(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}
