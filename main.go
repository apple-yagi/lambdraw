package main

import (
	"github.com/apple-yagi/lambdraw/config"
	"github.com/apple-yagi/lambdraw/pkg/handler"
	"github.com/apple-yagi/lambdraw/pkg/resizer"
	"github.com/apple-yagi/lambdraw/pkg/s3"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	conf := config.NewAwsConfig()
	client := s3.NewClient(conf)
	resizer := resizer.NewResizer()
	handler := handler.NewHandler(client, resizer)
	lambda.Start(handler.Execute)
}
