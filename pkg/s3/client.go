package s3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	cfg "github.com/igefined/nftique/pkg/config"
)

//go:generate mockgen -source=client.go -package=mocks -destination=./mocks/mock_s3.go S3

type S3 interface {
	List(ctx context.Context, filename string) ([]*Media, error)
	Store(ctx context.Context, filename string, contentBytes []byte) error
}

type Client struct {
	awsCfg   *cfg.AWSCfg
	client   *s3.S3
	uploader *s3manager.Uploader

	bucketName string
}

//nolint:ireturn
func New(awsCfg *cfg.AWSCfg, opts ...Opt) (*Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(awsCfg.AWSAccessKey, awsCfg.AWSSecretKey, ""),
		Region:      aws.String(awsCfg.AWSRegion)},
	)
	if err != nil {
		return nil, err
	}

	client := s3.New(sess)

	instance := &Client{
		awsCfg: awsCfg,
		client: client,
	}

	for _, opt := range opts {
		opt(instance)
	}

	return instance, nil
}
