package repository

import (
	"bytes"
	"fmt"
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

func (r S3Repository) Put(binary []byte) (string, error) {
	result, err := r.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("resizeapi"),
		Key: aws.String("gopher.png"),
		Body: bytes.NewReader(binary),
		ACL:    aws.String("public-read"),
	}); 
	if err != nil {
		fmt.Println("S3 upload error")
		return "", err
	}

	return result.Location ,nil
}
