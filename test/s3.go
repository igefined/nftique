package test

import (
	"context"
	"fmt"

	cfg "github.com/igefined/nftique/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
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

func NewS3Container(ctx context.Context, s3Cfg *cfg.S3, awsCfg *cfg.AWSCfg, opt *Opt) (*S3Container, error) {
	provider, err := testcontainers.NewDockerProvider()
	if err != nil {
		return nil, err
	}
	defer provider.Close()

	localstackContainer, err := localstack.RunContainer(
		ctx, testcontainers.WithImage(opt.Image))
	if err != nil {
		return nil, err
	}

	port, err := localstackContainer.MappedPort(ctx, nat.Port("4566/tcp"))
	if err != nil {
		return nil, err
	}

	host, err := provider.DaemonHost(ctx)
	if err != nil {
		return nil, err
	}

	return &S3Container{
		s3Cfg:               s3Cfg,
		awsCfg:              awsCfg,
		LocalStackContainer: localstackContainer,
		endpoint:            fmt.Sprintf("http://%s:%d", host, port.Int()),
	}, nil
}

func (c *S3Container) S3Client(ctx context.Context) (*s3.Client, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, opts ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           c.endpoint,
				SigningRegion: region,
			}, nil
		})

	options, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithRegion(c.awsCfg.AWSRegion),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(c.awsCfg.AWSAccessKeyID, c.awsCfg.AWSSecretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(options)

	return s3Client, nil
}
