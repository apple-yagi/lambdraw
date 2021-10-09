package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"resize-api/pkg/resizer"
	"resize-api/pkg/s3"

	"github.com/aws/aws-lambda-go/events"
)

type Request events.APIGatewayProxyRequest

type Response events.APIGatewayProxyResponse

type Handler struct {
	Client s3.S3
	Resizer *resizer.Resizer
}

func NewHandler(c s3.S3, r *resizer.Resizer) *Handler {
	return &Handler{Client: c, Resizer: r}
}

func (h *Handler) Execute(req Request) *Response {
	if len(req.Body) == 0 {
		return h.newResponse("", errors.New("must request body"))
	}

	data, err := base64.StdEncoding.DecodeString(req.Body)
	if err != nil {
		return h.newResponse("", errors.New("failed DecodeString req.Body"))
	}

	buff, err := h.Resizer.Resize(data);
	if err != nil {
		return h.newResponse("", errors.New("failed Resize"))
	}

	url, err := h.Client.PutImage("gopher.png", buff);
	if err != nil {
		return h.newResponse("", errors.New("failed PutImage"))
	}

	body, err := json.Marshal(map[string]interface{}{
		"url": url,
	})
	if err != nil {
		return h.newResponse("", errors.New("failed Marshal"))
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	return h.newResponse(buf.String(), nil)
}

func (h *Handler) newResponse(body string, err error) *Response {
	if err != nil {
		return &Response{
			StatusCode: 500,
			Body: err.Error(),
		}
	}

	return &Response{
		StatusCode: 200,
		Body:       body,
		Headers: map[string]string{
			"Content-Type":           "application/json",
		},
	}
}