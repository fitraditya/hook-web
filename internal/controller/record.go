package controller

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"net/url"

	"github.com/fitraditya/hook-web/internal/constant"
	"github.com/fitraditya/hook-web/internal/model"
	"github.com/fitraditya/hook-web/internal/response"
	"github.com/fitraditya/hook-web/internal/schema"
	"github.com/go-chi/chi/v5"
	"github.com/r3labs/sse/v2"
)

type RecordController struct {
	request model.Request
	sse     *sse.Server
}

func NewRecordController(request model.Request, sse *sse.Server) RecordController {
	return RecordController{
		request: request,
		sse:     sse,
	}
}

func (rec RecordController) Save(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	requestInfo := &schema.RequestInfo{
		Method:        r.Method,
		Path:          r.RequestURI,
		IP:            r.RemoteAddr,
		ContentType:   r.Header.Get("content-type"),
		ContentLength: r.ContentLength,
	}

	headers := schema.M{}
	queries := schema.M{}
	body, _ := io.ReadAll(r.Body)

	for k := range r.Header {
		headers[k] = r.Header.Get(k)
	}

	for k, v := range r.URL.Query() {
		if len(v) > 1 {
			queries[k] = v
		} else {
			queries[k] = v[0]
		}
	}

	requestInfo.Headers = headers
	requestInfo.Query = queries
	requestInfo.Data = string(body)

	contentType := r.Header.Get(constant.HeaderContentType)
	mediaType, _, _ := mime.ParseMediaType(contentType)

	if mediaType == constant.MIMEApplicationJSON {
		var jsonBody schema.M
		_ = json.Unmarshal(body, &jsonBody)
		requestInfo.Body = jsonBody
	} else if mediaType == constant.MIMEApplicationForm {
		forms := schema.M{}
		formValues, _ := url.ParseQuery(string(body))

		for k, v := range formValues {
			if len(v) > 1 {
				forms[k] = v
			} else {
				forms[k] = v[0]
			}
		}

		requestInfo.Form = forms
	}

	slug := chi.URLParam(r, "slug")

	record, err := rec.request.Create(ctx, slug, requestInfo)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(record)
	if err != nil {
		panic(err)
	}

	rec.sse.Publish(slug, &sse.Event{
		ID:   []byte(record.ID.Hex()),
		Data: data,
	})

	response := response.RecordResponse{
		Success: true,
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
