package s3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	cfg "github.com/igefined/nftique/pkg/config"
)

type Client struct {
	awsCfg *cfg.AWSCfg
	client *s3.S3

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

func (c *Client) Get(ctx context.Context, name string) (*Media, error) {
	return nil, nil
}
