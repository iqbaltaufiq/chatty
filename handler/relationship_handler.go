package handler

import "net/http"

type RelationshipHandler interface {
	Befriend(w http.ResponseWriter, r *http.Request)
	Unfriend(w http.ResponseWriter, r *http.Request)
	UnfriendByID(w http.ResponseWriter, r *http.Request)
	GetFriends(w http.ResponseWriter, r *http.Request)
}
