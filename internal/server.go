package internal

import (
	"net/http"
)

func NewHttpServer(addr string, db DB) *http.Server {
	handler := NewUrlShortener(db)

	r := &http.ServeMux{}
	r.HandleFunc("GET /{id}", handler.HandleGet)
	r.HandleFunc("POST /", handler.HandlePost)

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
