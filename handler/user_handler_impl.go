package handler

import (
	"net/http"

	"chatty/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService}
}

// Get retrieves user's detail identified by username.
// Password field is excluded from the response.
func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		JSONError(w, http.StatusBadRequest, &PathErrNoPathValue{})
		return
	}
	user, err := h.userService.GetUser(r.Context(), username)
	if err != nil {
		JSONError(w, http.StatusInternalServerError, err)
		return
	}
	JSONSuccess(w, http.StatusOK, M{
		"user": user,
	})
}

// GetByUUID retrieves user's detail identified by userID in UUID.
// Password field is excluded from the response.
func (h *userHandler) GetByUUID(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}

// UpdateFullname only update `fullname` field on the database
// instead of updating the whole row.
func (h *userHandler) UpdateFullname(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}

// UpdatePhoto only update `photo` field on the database
// instead of updating the whole row.
//
// Allowed file extensions: jpg, jpeg, png, webp, avif, tiff.
func (h *userHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	// not yet implemented
}
