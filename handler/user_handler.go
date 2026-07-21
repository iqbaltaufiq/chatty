package handler

import "net/http"

type UserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetCurrent(w http.ResponseWriter, r *http.Request)
	UpdateFullname(w http.ResponseWriter, r *http.Request)
	UpdatePhoto(w http.ResponseWriter, r *http.Request)
}
