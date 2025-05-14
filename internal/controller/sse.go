package controller

import (
	"net/http"

	"github.com/r3labs/sse/v2"
)

type SSEController struct {
	server *sse.Server
}

func NewSSEController(server *sse.Server) SSEController {
	return SSEController{
		server: server,
	}
}

func (sse SSEController) Serve(w http.ResponseWriter, r *http.Request) {
	sse.server.ServeHTTP(w, r)
}
