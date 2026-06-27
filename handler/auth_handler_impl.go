package handler

import (
	"encoding/json"
	"net/http"

	"chatty/entity"
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

// Register handles account creation for new user.
func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}
	if err := h.authService.CreateUser(r.Context(), user); err != nil {
		jsonError(w, http.StatusInternalServerError, err)
		return
	}
	payload := M{"message": "your account has been created"}
	jsonSuccess(w, http.StatusCreated, payload)
}
