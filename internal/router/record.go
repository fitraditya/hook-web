package router

import (
	"github.com/fitraditya/hook-web/internal/controller"
	"github.com/fitraditya/hook-web/internal/model"
	"github.com/fitraditya/hook-web/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/r3labs/sse/v2"
)

func recordRoutes(sse *sse.Server, mongo *service.MongoDB) chi.Router {
	res := controller.NewRecordController(
		*model.NewRequest(mongo),
		sse,
	)

	r := chi.NewRouter()
	r.HandleFunc("/", res.Save)

	return r
}
