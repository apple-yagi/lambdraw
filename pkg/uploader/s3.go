package uploader

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Put() error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))

	uploader := s3manager.NewUploader(sess)
	
	f, err := os.Open("./tmp/original/gopher.png")
	if err != nil {
		return err
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("resizeapi"),
		Key: aws.String("test"),
		Body: f,
	})
	if err != nil {
		return err
	}

	return nil
}
