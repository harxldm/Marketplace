package handler

import "net/http"

type storage interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
