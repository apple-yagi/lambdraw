package uploader

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Uploader struct {
	Uploader *s3manager.Uploader
}

func NewS3Uploader() Uploader {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))

	uploader := s3manager.NewUploader(sess)

	return &S3Uploader{Uploader: uploader}
}

func (s3 *S3Uploader) Execute() error {
	f, err := os.Open("./tmp/original/gopher.png")
	if err != nil {
		return err
	}

	_, err = s3.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("resizeapi"),
		Key: aws.String("test"),
		Body: f,
	})
	if err != nil {
		return err
	}

	return nil
}
