package s3

import (
	"context"

	cfg "github.com/igefined/nftique/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:generate mockgen -source=client.go -package=mocks -destination=./mocks/mock_s3.go S3

type S3 interface {
	List(ctx context.Context, filename string) ([]*Media, error)
	Store(ctx context.Context, filename string, contentBytes []byte) error
}

type Client struct {
	awsCfg   *cfg.AWSCfg
	client   *s3.Client
	uploader *manager.Uploader

	bucketName string
}

//nolint:ireturn
func New(awsCfg *cfg.AWSCfg, opts ...Opt) (*Client, error) {
	options := s3.Options{
		Region: awsCfg.AWSRegion,
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(awsCfg.AWSAccessKeyID, awsCfg.AWSSecretKey, "")),
	}

	s3Client := s3.New(options)

	instance := &Client{
		awsCfg:   awsCfg,
		client:   s3Client,
		uploader: manager.NewUploader(s3Client),
	}

	for _, opt := range opts {
		opt(instance)
	}

	return instance, nil
}
