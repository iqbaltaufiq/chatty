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
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	creds.Username = strings.TrimSpace(creds.Username)

	if err := h.authService.ValidateCreds(r.Context(), creds); err != nil {
		util.JSONError(w, http.StatusUnauthorized, err)
		return
	}
	user, err := h.userService.GetUser(r.Context(), creds.Username)
	if err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}

	// create auth cookie that contains jwt as its value.
	jsonwebtoken, err := util.GenerateJWT(user.UUID, "user")
	if err != nil {
		util.JSONError(w, http.StatusUnauthorized, err)
		return
	}
	authCookie := util.SetAuthCookie(jsonwebtoken, 60*60*24)
	http.SetCookie(w, authCookie)

	util.JSONSuccess(w, http.StatusOK, util.M{
		"message": "you are successfully logged in",
	})
}

// Register handles account creation for new user.
func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	const (
		maxFileSize int64 = 2 * 1000 * 1000 // 2MB
		maxMemory   int64 = 1 << 10         // 1KB
	)
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	file, header, err := r.FormFile("photo")
	if err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	if header.Size > maxFileSize {
		util.JSONError(w, http.StatusBadRequest, &RequestErrSizeLimitExceeded{})
		return
	}
	user := entity.User{
		Username: r.FormValue("username"),
		Fullname: r.FormValue("fullname"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Photo:    header.Filename,
	}
	if err := h.authService.CreateUser(r.Context(), user, file); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	payload := util.M{"message": "your account has been created"}
	util.JSONSuccess(w, http.StatusCreated, payload)
}

// Logout deauthenticates user by removing their auth cookie.
func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	deletedCookie := util.DeleteCookie(util.AuthCookieName)
	http.SetCookie(w, deletedCookie)
	util.JSONSuccess(w, http.StatusOK, util.M{"message": "logout successfully"})
}
