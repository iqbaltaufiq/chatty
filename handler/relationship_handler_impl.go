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
	friendshipID := r.PathValue("friendshipID")
	if friendshipID == "" {
		util.JSONError(w, http.StatusBadRequest, &PathErrNoPathValue{})
		return
	}
	i8FriendshipID, err := strconv.Atoi(friendshipID)
	if err != nil {
		util.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if err := h.relationshipService.RemoveByID(r.Context(), i8FriendshipID); err != nil {
		util.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	util.JSONSuccess(w, http.StatusOK, util.M{
		"message": "successfully removed a user from friend list",
	})
}

func (h *relationshipHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	// wip
}
