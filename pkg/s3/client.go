package s3

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"resize-api/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"golang.org/x/image/draw"
)

type Client struct {
	Uploader *s3manager.Uploader
	Conf *config.AwsConfig
}

func NewClient(conf *config.AwsConfig) *Client {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(conf.RegionName),
	}))

	uploader := s3manager.NewUploader(sess)

	return &Client{Uploader: uploader, Conf: conf}
}

func (c *Client) PutImage(binary []byte) (string, error) {
	reader := bytes.NewReader(binary)
	img, t, err := image.Decode(reader)
	if err != nil {
		log.Panic(err)
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

	result, err := c.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.Conf.BucketName),
		Key: aws.String("gopppher" + "." + t),
		Body: bytes.NewReader(buff.Bytes()),
		ACL:    aws.String("public-read"),
	}); 
	if err != nil {
		log.Panic(err)
		return "", err
	}

	return result.Location ,nil
}
