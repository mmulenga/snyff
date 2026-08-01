package api

import (
	"encoding/json"
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
	if _, err := io.WriteString(w, "Status: OK!\n"); err != nil {
		log.Println(err)
	}
}

func (r *Router) ListRequestHandler(w http.ResponseWriter, req *http.Request) {
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
	reqJSON, err := json.Marshal(requests)
	if err != nil {
		log.Println(err)
	}
	if _, err := fmt.Fprint(w, string(reqJSON)); err != nil {
		log.Println(err)
	}
}

func (r *Router) FindRequestHandler(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")

	request, err := r.requestStore.FindById(id)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := fmt.Fprintf(w, `%v`, *request); err != nil {
		log.Println(err)
	}
}
