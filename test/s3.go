package test

import (
	"context"
	"fmt"

	cfg "github.com/igefined/nftique/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

type S3Container struct {
	s3Cfg    *cfg.S3
	awsCfg   *cfg.AWSCfg
	endpoint string

	*localstack.LocalStackContainer
}

func NewS3Container(ctx context.Context, cfg *cfg.S3, awsCfg *cfg.AWSCfg, opt *Opt) (*S3Container, error) {
	localstackContainer, err := localstack.RunContainer(ctx, testcontainers.WithImage(opt.Image))
	if err != nil {
		return nil, err
	}

	port, err := localstackContainer.MappedPort(ctx, nat.Port("4566/tcp"))
	if err != nil {
		return nil, err
	}

	host, err := localstackContainer.Host(ctx)
	if err != nil {
		return nil, err
	}

	return &S3Container{
		s3Cfg:               cfg,
		awsCfg:              awsCfg,
		LocalStackContainer: localstackContainer,
		endpoint:            fmt.Sprintf("http://%s:%d", host, port.Int()),
	}, nil
}

func (c *S3Container) S3Client() (*s3.Client, error) {
	options := s3.Options{
		// Region:       c.awsCfg.AWSRegion,
		BaseEndpoint: aws.String(c.endpoint),
		UsePathStyle: true,
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(c.awsCfg.AWSAccessKey, c.awsCfg.AWSSecretKey, "")),
	}

	s3Client := s3.New(options)

	return s3Client, nil
}
