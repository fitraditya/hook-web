package response

import "github.com/fitraditya/hook-web/internal/schema"

type RecordResponse struct {
	Success bool `json:"success"`
}

type ApiResponse struct {
	Result []*schema.Request `json:"result"`
}
