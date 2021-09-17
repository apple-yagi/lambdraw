package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"resize-api/pkg/domain"

	"github.com/aws/aws-lambda-go/events"
)

// Request is of type APIGatewayProxyRequest
type Request events.APIGatewayProxyRequest

type JsonBody struct {
	Image  string `json:"image"`
	UserId int    `json:"user_id"`
}

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
	data, decodeErr := base64.StdEncoding.DecodeString(req.Body)
	if decodeErr != nil {
		fmt.Println(decodeErr)
		return Response{StatusCode: 500}, decodeErr
	}

	url, err := h.Repository.Put(data); 
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"url": url,
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