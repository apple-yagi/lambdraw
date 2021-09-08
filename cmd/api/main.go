package main

import (
	"resize-api/pkg/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handler)
}
