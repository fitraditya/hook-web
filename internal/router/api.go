package router

import (
	"github.com/fitraditya/hook-web/internal/controller"
	"github.com/fitraditya/hook-web/internal/model"
	"github.com/fitraditya/hook-web/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/r3labs/sse/v2"
)

func apiRoutes(sse *sse.Server, mongo *service.MongoDB) chi.Router {
	res := controller.NewApiController(
		*model.NewRequest(mongo),
		sse,
	)

	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		//r.Get("list", res.List)
		r.Get("/{slug:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}/inspect", res.Inspect)
	})

	return r
}
