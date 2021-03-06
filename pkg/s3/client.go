package s3

import (
	"bytes"

	"github.com/apple-yagi/lambdraw/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Client struct {
	S3

	Uploader *s3manager.Uploader
	Conf     *config.AwsConfig
}

type S3 interface {
	PutImage(key string, buff *bytes.Buffer) (string, error)
}

func NewClient(conf *config.AwsConfig) *Client {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(conf.RegionName),
	}))

	uploader := s3manager.NewUploader(sess)

	return &Client{Uploader: uploader, Conf: conf}
}

func (c *Client) PutImage(key string, buff *bytes.Buffer) (string, error) {
	result, err := c.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.Conf.BucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buff.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
