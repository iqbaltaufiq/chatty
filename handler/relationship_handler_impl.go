package handler

import (
	"net/http"
	"strconv"

	"chatty/service"
	"chatty/util"
)

type relationshipHandler struct {
	relationshipService service.RelationshipService
}

func NewRelationshipHandler(relationshipService service.RelationshipService) RelationshipHandler {
	return &relationshipHandler{relationshipService}
}

func (h *relationshipHandler) Befriend(w http.ResponseWriter, r *http.Request) {
	target := r.PathValue("username")
	if target == "" {
		util.JSONError(w, http.StatusBadRequest, &PathErrNoPathValue{})
		return
	}
	self, _ := util.FromCtxUsername(r.Context())
	if err := h.relationshipService.Add(r.Context(), self, target); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"message": "successfully befriending a user",
	})
}

func (h *relationshipHandler) Unfriend(w http.ResponseWriter, r *http.Request) {
	target := r.PathValue("username")
	if target == "" {
		util.JSONError(w, http.StatusBadRequest, &PathErrNoPathValue{})
		return
	}
	self, _ := util.FromCtxUsername(r.Context())
	if err := h.relationshipService.Remove(r.Context(), self, target); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"message": "successfully removed a user from friend list",
	})
}

func (h *relationshipHandler) UnfriendByID(w http.ResponseWriter, r *http.Request) {
	relationshipID := r.PathValue("relationshipID")
	if relationshipID == "" {
		util.JSONError(w, http.StatusBadRequest, &PathErrNoPathValue{})
		return
	}
	i8relationshipID, err := strconv.Atoi(relationshipID)
	if err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if err := h.relationshipService.RemoveByID(r.Context(), i8relationshipID); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"message": "successfully removed a user from friend list",
	})
}

func (h *relationshipHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	username, _ := util.FromCtxUsername(r.Context())
	friends, err := h.relationshipService.GetFriends(r.Context(), username)
	if err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"message": "found",
		"friends": friends,
	})
}
