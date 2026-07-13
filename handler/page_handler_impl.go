package handler

import (
	"net/http"

	"chatty/service"
)

type pageHandler struct {
	userService service.UserService
}

func NewPageHandler(userService service.UserService) PageHandler {
	return &pageHandler{userService}
}

func (h *pageHandler) Index(w http.ResponseWriter, r *http.Request) {
	// wip
}

func (h *pageHandler) Login(w http.ResponseWriter, r *http.Request) {
	// wip
}

func (h *pageHandler) Register(w http.ResponseWriter, r *http.Request) {
	// wip
}
