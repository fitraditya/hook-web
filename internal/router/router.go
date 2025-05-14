package router

import (
	"github.com/fitraditya/hook-web/internal/service"
	"github.com/go-chi/chi/v5"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) ApiRoutes() chi.Router {
	sse := service.NewSSEServer()
	mongo := service.NewMongoDB()

	router := chi.NewRouter()
	router.Mount("/sse", sseRoutes(sse))
	router.Mount("/api", apiRoutes(sse, mongo))
	router.Mount("/{slug:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", recordRoutes(sse, mongo))

	return router
}

func (r *Router) WebRoutes() chi.Router {
	router := chi.NewRouter()
	router.Mount("/", webRoutes())

	return router
}
