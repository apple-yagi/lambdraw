package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"resize-api/pkg/s3"

	"github.com/aws/aws-lambda-go/events"
)

// Request is of type APIGatewayProxyRequest
type Request events.APIGatewayProxyRequest

// Response is of type APIGatewayProxyResponse
type Response events.APIGatewayProxyResponse

type Handler struct {
	Client *s3.Client
}

func NewHandler(c *s3.Client) *Handler {
	return &Handler{Client: c}
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func (h *Handler) Execute(req Request) (Response, error) {
	data, decodeErr := base64.StdEncoding.DecodeString(req.Body)
	if decodeErr != nil {
		log.Panic(decodeErr)
		return Response{StatusCode: 500}, decodeErr
	}

	url, err := h.Client.PutImage(data); 
	if err != nil {
		log.Panic(err)
		return Response{StatusCode: 500}, err
	}

	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"url": url,
	})
	if err != nil {
		log.Panic(err)
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "world-handler",
		},
	}

	return resp, nil
}