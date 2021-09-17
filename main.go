package main

import (
	"resize-api/config"
	"resize-api/pkg/handler"
	"resize-api/pkg/repository"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	conf := config.NewAwsConfig()
	repository := repository.NewS3Repository(conf)
	handler := handler.NewHandler(repository)
	lambda.Start(handler.Execute)
}
