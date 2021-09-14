package main

import (
	"resize-api/pkg/handler"
	"resize-api/pkg/repository"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	repository := repository.NewS3Repository()
	handler := handler.NewHandler(repository)
	lambda.Start(handler.Execute)
}
