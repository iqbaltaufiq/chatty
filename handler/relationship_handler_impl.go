package handler

import (
	"net/http"

	"chatty/service"
)

type relationshipHandler struct {
	relationshipService service.RelationshipService
}

func NewRelationshipHandler(relationshipService service.RelationshipService) RelationshipHandler {
	return &relationshipHandler{relationshipService}
}

func (h *relationshipHandler) Befriend(w http.ResponseWriter, r *http.Request) {
	// wip
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
