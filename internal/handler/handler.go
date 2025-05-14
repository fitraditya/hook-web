package handler

import (
	"net/http"
	"time"

	"github.com/fitraditya/hook-web/config"
	"github.com/fitraditya/hook-web/internal/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ApiServer() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	if config.EnableRateLimit() {
		r.Use(httprate.Limit(
			config.GetRateLimit(),
			time.Second,
			httprate.WithKeyFuncs(httprate.KeyByEndpoint),
		))
	}

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Ok"))
	})

	r.Mount("/", router.NewRouter().ApiRoutes())

	return r
}

func (h *Handler) WebServer() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Mount("/", router.NewRouter().WebRoutes())

	return r
}
