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
		JSONError(w, http.StatusBadRequest, &PathErrNoPathValue{})
		return
	case usernameLimit:
		user, err = h.userService.GetUser(r.Context(), identifier)
		if err != nil {
			JSONError(w, http.StatusInternalServerError, err)
			return
		}
	case uuidLimit:
		_, err = util.UUIDtoBinary(identifier)
		if err != nil {
			JSONError(w, http.StatusBadRequest, err)
			return
		}
		user, err = h.userService.GetUserByUUID(r.Context(), identifier)
		if err != nil {
			JSONError(w, http.StatusInternalServerError, err)
			return
		}
	default:
		JSONError(w, http.StatusInternalServerError, &PathErrInvalidPath{})
		return
	}
	JSONSuccess(w, http.StatusOK, M{
		"user": user,
	})
}

// UpdateFullname only update `fullname` field on the database
// instead of updating the whole row.
func (h *userHandler) UpdateFullname(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Fullname string `json:"fullname"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}
	if payload.Fullname == "" {
		JSONError(w, http.StatusBadRequest, &RequestErrInsufficientBody{})
		return
	}
	username, ok := util.FromCtxUsername(r.Context())
	if !ok {
		JSONError(w, http.StatusUnauthorized, &util.ContextErrNotFound{})
		return
	}
	if err := h.userService.UpdateFullname(
		r.Context(),
		username,
		payload.Fullname,
	); err != nil {
		JSONError(w, http.StatusInternalServerError, err)
		return
	}
	JSONSuccess(w, http.StatusOK, M{
		"message": "fullname has been updated successfully",
	})
}

// UpdatePhoto only update `photo` field on the database
// instead of updating the whole row.
//
// Allowed file extensions: jpg, jpeg, png, webp, avif, tiff.
func (h *userHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}
