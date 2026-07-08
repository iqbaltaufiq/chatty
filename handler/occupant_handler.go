package handler

import "net/http"

type OccupantHandler interface {
	Add(w http.ResponseWriter, r *http.Request)
	Remove(w http.ResponseWriter, r *http.Request)
}
