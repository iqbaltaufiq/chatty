package handler

import (
	"encoding/json"
	"net/http"

	"chatty/entity"
	"chatty/model"
	customerrors "chatty/model/custom_errors"
	"chatty/service"
	"chatty/util"
)

type userHandler struct {
	userService service.UserService
	fileService service.FileService
}

func NewUserHandler(
	userService service.UserService,
	fileService service.FileService,
) UserHandler {
	return &userHandler{userService, fileService}
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
	util.SendJSON(w, http.StatusOK, userPayload{user})
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
	util.SendJSON(w, http.StatusOK, userPayload{user})
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
	util.SendJSON(w, http.StatusOK, messagePayload{"fullname has been updated successfully"})
}

// UpdatePhoto only update `photo` field on the database
// instead of updating the whole row.
//
// Allowed file extensions: jpg, jpeg, png, webp, avif, tiff.
func (h *userHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	const maxMemory int64 = 1 << 10 // 1KB

	if err := r.ParseMultipartForm(maxMemory); err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}

	oldPhoto := r.FormValue("old_photo")
	if oldPhoto == "" {
		util.JSONError(w, http.StatusBadRequest, &customerrors.ErrInsufficientPayload{})
		return
	}
	file, header, err := r.FormFile("photo")
	if err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	username, ok := util.FromCtxUsername(r.Context())
	if !ok {
		util.JSONError(w, http.StatusUnauthorized, &util.ContextErrNotFound{})
		return
	}
	up := model.UploadedPhoto{
		File:   file,
		Header: header,
	}
	if err := up.IsValid(); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	hashedFilename := up.HashFilename()
	err = h.userService.UpdatePhoto(r.Context(), username, hashedFilename)
	if err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	err = h.fileService.StoreAvatar(r.Context(), file, hashedFilename)
	if err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	if oldPhoto != service.DefaultPhoto {
		if err := h.fileService.Delete(r.Context(), oldPhoto); err != nil {
			util.JSONError(w, http.StatusInternalServerError, err)
			return
		}
	}
	util.SendJSON(w, http.StatusOK, messagePayload{"photo has been updated successfully"})
}
