package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"snyff/internal/store"
	"strconv"
)

type Router struct {
	requestStore store.RequestRepository
}

func NewRouter(s store.RequestRepository) *Router {
	return &Router{s}
}

func (r *Router) HealthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Status: OK!\n")
}

func (r *Router) RequestHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 10 {
		limit = 10
	}

	offset := (page - 1) * limit
	requests, err := r.requestStore.List(offset, limit)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"page":%d,"limit":%d,"items":%v}`, page, limit, requests)
}
