package handler

import (
	"net/http"

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
	// wip
}

func (h *relationshipHandler) UnfriendByID(w http.ResponseWriter, r *http.Request) {
	// wip
}

func (h *relationshipHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	// wip
}
