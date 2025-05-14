package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func webRoutes() chi.Router {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("web/dist"))
	r.Handle("/*", http.StripPrefix("/", fs))

	return r
}
