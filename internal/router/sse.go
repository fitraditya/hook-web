package router

import (
	"github.com/fitraditya/hook-web/internal/controller"
	"github.com/go-chi/chi/v5"
	"github.com/r3labs/sse/v2"
)

func sseRoutes(sse *sse.Server) chi.Router {
	res := controller.NewSSEController(sse)

	r := chi.NewRouter()
	r.HandleFunc("/", res.Serve)

	return r
}
