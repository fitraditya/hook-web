package controller

import (
	"encoding/json"
	"net/http"

	"github.com/fitraditya/hook-web/internal/model"
	"github.com/fitraditya/hook-web/internal/response"
	"github.com/go-chi/chi/v5"
	"github.com/r3labs/sse/v2"
)

type ApiController struct {
	request model.Request
	sse     *sse.Server
}

func NewApiController(request model.Request, sse *sse.Server) ApiController {
	return ApiController{
		request: request,
		sse:     sse,
	}
}

func (api ApiController) Inspect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := chi.URLParam(r, "slug")

	requests, err := api.request.List(ctx, slug)
	if err != nil {
		panic(err)
	}

	response := response.Response{
		Result: requests,
	}

	data, _ := json.Marshal(response)
	w.Write(data)
}
