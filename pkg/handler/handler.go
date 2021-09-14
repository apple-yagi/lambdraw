package handler

import (
	"bytes"
	"encoding/json"
	"resize-api/pkg/domain"

	"github.com/aws/aws-lambda-go/events"
)

// Request is of type APIGatewayProxyRequest
type Request events.APIGatewayProxyRequest

// Response is of type APIGatewayProxyResponse
type Response events.APIGatewayProxyResponse

type Handler struct {
	Repository domain.Repository
}

func NewHandler(r domain.Repository) Handler {
	return Handler{Repository: r}
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func (h *Handler) Execute(req Request) (Response, error) {
	err := h.Repository.Put(); if err != nil {
		return Response{StatusCode: 500}, err
	}

	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Okay so your other function also executed successfully!",
	})
	if err != nil {
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