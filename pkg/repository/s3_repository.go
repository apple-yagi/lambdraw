package repository

import (
	"os"
	"resize-api/pkg/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Repository struct {
	Uploader *s3manager.Uploader
}

func NewS3Repository() domain.Repository {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))

	uploader := s3manager.NewUploader(sess)

	return S3Repository{Uploader: uploader}
}

func (r S3Repository) Put() error {
	f, err := os.Open("./tmp/original/gopher.png")
	if err != nil {
		return err
	}

	_, err = r.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("resizeapi"),
		Key: aws.String("test"),
		Body: f,
	})
	if err != nil {
		return err
	}

	return nil
}
