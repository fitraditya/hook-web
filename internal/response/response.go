package response

import "github.com/fitraditya/hook-web/internal/schema"

type Response struct {
	Result []*schema.Request `json:"result"`
}
