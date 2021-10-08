package main

import (
	"resize-api/config"
	"resize-api/pkg/handler"
	"resize-api/pkg/resizer"
	"resize-api/pkg/s3"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	conf := config.NewAwsConfig()
	client := s3.NewClient(conf)
	resizer := resizer.NewResizer()
	handler := handler.NewHandler(client, resizer)
	lambda.Start(handler.Execute)
}
