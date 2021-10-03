package main

import (
	"resize-api/config"
	"resize-api/pkg/handler"
	"resize-api/pkg/s3"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	conf := config.NewAwsConfig()
	client := s3.NewClient(conf)
	handler := handler.NewHandler(client)
	lambda.Start(handler.Execute)
}
