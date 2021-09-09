package main

import (
	"resize-api/pkg/handler"
	"resize-api/pkg/uploader"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	uploader := uploader.NewS3Uploader()
	handler := handler.NewHandler(uploader)
	lambda.Start(handler.Execute)
}
