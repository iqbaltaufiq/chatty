package handler

import (
	"encoding/json"
	"net/http"

	"chatty/entity"
	"chatty/service"
	"chatty/util"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService}
}

// Get retrieves user's detail identified by username or uuid.
// Password field is excluded from the response.
//
// Username has a maximum length of 32,
// meanwhile a uuid string has a fixed length of 36.
// Request is rejected when the path value does not meet either expectation.
func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	var (
		user entity.User
		err  error

		identifier    = r.PathValue("identifier")
		identifierLen = len(identifier)
		usernameLimit = identifierLen > 0 && identifierLen <= 32
		uuidLimit     = identifierLen == 36
	)

	switch {
	case identifier == "":
		util.JSONError(w, http.StatusBadRequest, &PathErrNoPathValue{})
		return
	case usernameLimit:
		user, err = h.userService.GetUser(r.Context(), identifier)
		if err != nil {
			util.JSONError(w, http.StatusInternalServerError, err)
			return
		}
	case uuidLimit:
		_, err = util.UUIDtoBinary(identifier)
		if err != nil {
			util.JSONError(w, http.StatusBadRequest, err)
			return
		}
		user, err = h.userService.GetUserByUUID(r.Context(), identifier)
		if err != nil {
			util.JSONError(w, http.StatusInternalServerError, err)
			return
		}
	default:
		util.JSONError(w, http.StatusInternalServerError, &PathErrInvalidPath{})
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"user": user,
	})
}

// GetCurrent retrieves information about the currently logged in user.
// It gets the username from the request context.
func (h *userHandler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	username, _ := util.FromCtxUsername(r.Context())

	user, err := h.userService.GetUser(r.Context(), username)
	if err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{"user": user})
}

// UpdateFullname only update `fullname` field on the database
// instead of updating the whole row.
func (h *userHandler) UpdateFullname(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Fullname string `json:"fullname"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if payload.Fullname == "" {
		util.JSONError(w, http.StatusBadRequest, &RequestErrInsufficientBody{})
		return
	}
	username, ok := util.FromCtxUsername(r.Context())
	if !ok {
		util.JSONError(w, http.StatusUnauthorized, &util.ContextErrNotFound{})
		return
	}
	if err := h.userService.UpdateFullname(
		r.Context(),
		username,
		payload.Fullname,
	); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"message": "fullname has been updated successfully",
	})
}

// UpdatePhoto only update `photo` field on the database
// instead of updating the whole row.
//
// Allowed file extensions: jpg, jpeg, png, webp, avif, tiff.
func (h *userHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Photo string `json:"photo"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if payload.Photo == "" {
		util.JSONError(w, http.StatusBadRequest, &RequestErrInsufficientBody{})
		return
	}
	username, ok := util.FromCtxUsername(r.Context())
	if !ok {
		util.JSONError(w, http.StatusUnauthorized, &util.ContextErrNotFound{})
		return
	}
	if err := h.userService.UpdatePhoto(
		r.Context(),
		username,
		payload.Photo,
	); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"message": "photo has been updated successfully",
	})
}
