package api

import (
	"io"
	"net/http"
	"snyff/internal/store"
)

type Router struct {
	store store.RequestRepository
}

func NewRouter(s store.RequestRepository) *Router {
	return &Router{s}
}

func (i *Router) HealthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Status: OK!\n")
}
