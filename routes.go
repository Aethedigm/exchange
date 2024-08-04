package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {

	})

	fileServer := http.FileServer(http.Dir("./dist"))
	r.Handle("/*", fileServer)

	return r
}
