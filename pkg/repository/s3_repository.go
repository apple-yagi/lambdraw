package repository

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"resize-api/config"
	"resize-api/pkg/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"golang.org/x/image/draw"
)

type S3Repository struct {
	Uploader *s3manager.Uploader
	Conf *config.AwsConfig
}

func NewS3Repository(conf *config.AwsConfig) domain.Repository {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(conf.RegionName),
	}))

	uploader := s3manager.NewUploader(sess)

	return S3Repository{Uploader: uploader, Conf: conf}
}

func (r S3Repository) Put(binary []byte) (string, error) {
	reader := bytes.NewReader(binary)
	img, t, err := image.Decode(reader)
	if err != nil {
		fmt.Println("Image decode error")
		return "", err
	}

	rct := img.Bounds()
	canvas := image.NewRGBA(image.Rect(0, 0, rct.Dx()*2, rct.Dy()*2))
	draw.CatmullRom.Scale(canvas, canvas.Bounds(), img, rct, draw.Over, nil)

	buff := bytes.NewBuffer([]byte{})

	switch t {
		case "jpeg":
			jpeg.Encode(buff, canvas, &jpeg.Options{Quality: 95})
		default:
			png.Encode(buff, canvas)
	}

	result, err := r.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(r.Conf.BucketName),
		Key: aws.String("gopppher" + "." + t),
		Body: bytes.NewReader(buff.Bytes()),
		ACL:    aws.String("public-read"),
	}); 
	if err != nil {
		fmt.Println("S3 upload error")
		return "", err
	}

	return result.Location ,nil
}
