package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"chatty/entity"
	"chatty/model"
	"chatty/service"
	"chatty/util"
)

type authHandler struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthHandler(
	authService service.AuthService,
	userService service.UserService,
) AuthHandler {
	return &authHandler{
		authService,
		userService,
	}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds model.Credential

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}
	creds.Username = strings.TrimSpace(creds.Username)

	if err := h.authService.ValidateCreds(r.Context(), creds); err != nil {
		JSONError(w, http.StatusUnauthorized, err)
		return
	}
	user, _ := h.userService.GetUser(r.Context(), creds.Username)

	// create auth cookie that contains jwt as its value.
	jsonwebtoken, err := util.GenerateJWT(user.UUID, "user")
	if err != nil {
		JSONError(w, http.StatusUnauthorized, err)
		return
	}
	authCookie := util.SetAuthCookie(jsonwebtoken, 60*60*24)
	http.SetCookie(w, authCookie)

	JSONSuccess(w, http.StatusOK, M{
		"message": "you are successfully logged in",
	})
}

// Register handles account creation for new user.
func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}
	if err := h.authService.CreateUser(r.Context(), user); err != nil {
		JSONError(w, http.StatusInternalServerError, err)
		return
	}
	payload := M{"message": "your account has been created"}
	JSONSuccess(w, http.StatusCreated, payload)
}

// Logout deauthenticates user by removing their auth cookie.
func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	deletedCookie := util.DeleteCookie(util.AuthCookieName)
	http.SetCookie(w, deletedCookie)
	JSONSuccess(w, http.StatusOK, M{"message": "logout successfully"})
}
