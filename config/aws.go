package config

import (
	"os"
)

type AwsConfig struct {
	RegionName       string
	BucketName       string
}

func NewAwsConfig() *AwsConfig {
		return &AwsConfig{
			RegionName: os.Getenv("REGION_NAME"),
			BucketName: os.Getenv("BUCKET_NAME"),
		}
}
