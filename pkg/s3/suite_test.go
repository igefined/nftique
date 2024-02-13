package s3

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/igefined/nftique/pkg/config"
	"github.com/igefined/nftique/pkg/log"
	"github.com/igefined/nftique/test"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

const defaultImage = "localstack/localstack:1.4.0"

type Suite struct {
	suite.Suite
	ctx context.Context

	s3Cfg     *config.S3
	awsCfg    *config.AWSCfg
	logger    *zap.Logger
	container *test.S3Container

	client S3
}

func TestBuilderSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	logger, err := log.NewLogger(zap.DebugLevel)
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	_ = cancel

	type testConfig struct {
		sync.RWMutex
		config.S3     `mapstructure:",squash"`
		config.AWSCfg `mapstructure:",squash"`
	}

	var cfg *testConfig
	s.Require().NoError(config.GetConfig(&cfg, []*config.EnvVar{}))
	s.Require().NotEmpty(cfg.NFTBucketName)

	s.logger = logger
	s.s3Cfg = &cfg.S3
	s.awsCfg = &cfg.AWSCfg
	s.ctx = ctx

	s3Container, err := test.NewS3Container(ctx, s.s3Cfg, s.awsCfg, &test.Opt{Enabled: true, Image: defaultImage})
	s.Require().NoError(err)

	s3Client, err := s3Container.S3Client()
	s.Require().NoError(err)

	bucket, err := s3Client.CreateBucket(ctx, &s3.CreateBucketInput{Bucket: aws.String(s.s3Cfg.NFTBucketName)})
	s.Require().NoError(err)
	s.Require().NotNil(bucket)

	s.container = s3Container

	client, err := New(s.awsCfg, WithBucketName(s.s3Cfg.NFTBucketName))
	s.Require().NoError(err)
	s.client = client
}

func (s *Suite) TearDownSuite() {
	s3Client, err := s.container.S3Client()
	s.Require().NoError(err)

	objects, err := s3Client.ListObjectsV2(s.ctx, &s3.ListObjectsV2Input{Bucket: aws.String(s.s3Cfg.NFTBucketName)})
	s.Require().NoError(err)
	s.Require().NotNil(objects)

	toDel := make([]types.ObjectIdentifier, 0, len(objects.Contents))
	for _, obj := range objects.Contents {
		if obj.Key != nil {
			toDel = append(toDel, types.ObjectIdentifier{
				Key: aws.String(*obj.Key),
			})
		}
	}

	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(s.s3Cfg.NFTBucketName),
		Delete: &types.Delete{
			Objects: toDel,
		},
	}

	_, err = s3Client.DeleteObjects(s.ctx, input)
	s.Require().NoError(err)

	_, err = s3Client.DeleteBucket(s.ctx, &s3.DeleteBucketInput{Bucket: aws.String(s.s3Cfg.NFTBucketName)})
	s.Require().NoError(err)

	if err = s.container.Terminate(s.ctx); err != nil {
		s.logger.Error("error terminating s3 test container", zap.Error(err))
	}
}
