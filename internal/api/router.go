package api

import (
	"net/http"
	"snyff/internal/store"
)

type Router struct {
	store store.RequestRepository
}

func (r *Router) DefaultHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
