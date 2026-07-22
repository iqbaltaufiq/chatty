package handler

import "net/http"

type ChatHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetRecents(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
