package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/apple-yagi/lambdraw/pkg/resizer"
	"github.com/apple-yagi/lambdraw/pkg/s3"
	"github.com/aws/aws-lambda-go/events"
)

type Request events.APIGatewayProxyRequest

type Response events.APIGatewayProxyResponse

type Handler struct {
	Client  s3.S3
	Resizer *resizer.Resizer
}

func NewHandler(c s3.S3, r *resizer.Resizer) *Handler {
	return &Handler{Client: c, Resizer: r}
}

func (h *Handler) Execute(req Request) (*Response, error) {
	if len(req.Body) == 0 {
		return nil, errors.New("empty body")
	}

	data, err := base64.StdEncoding.DecodeString(req.Body)
	if err != nil {
		return nil, err
	}

	buff, err := h.Resizer.Resize(data)
	if err != nil {
		return nil, err
	}

	url, err := h.Client.PutImage("gopher.png", buff)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"url": url,
	})
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	return h.newResponse(buf.String()), err
}

func (h *Handler) newResponse(body string) *Response {
	return &Response{
		StatusCode: 200,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
