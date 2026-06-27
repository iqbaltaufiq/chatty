package handler

import (
	"net/http"

	"chatty/service"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{
		authService,
	}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}
