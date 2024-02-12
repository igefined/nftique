package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const urlMedia = "https://%s.s3.amazonaws.com/%s"

func (c *Client) List(ctx context.Context, filename string) ([]*Media, error) {
	objects, err := c.client.ListObjectsWithContext(ctx, &s3.ListObjectsInput{
		Bucket: aws.String(c.bucketName),
		Prefix: aws.String(filename),
	})
	if err != nil {
		return nil, err
	}

	if len(objects.Contents) == 0 {
		return nil, ErrNoContents
	}

	out := make([]*Media, 0, len(objects.Contents))

	for i := range objects.Contents {
		object := objects.Contents[i]
		if objects.Contents[i] == nil {
			continue
		}

		if objects != nil && *object.Size != 0 {
			u, _ := url.Parse(fmt.Sprintf(urlMedia, c.bucketName, *object.Key))
			media := &Media{
				Filename:     strings.Split(*object.Key, "/")[2],
				Date:         strings.Split(*object.Key, "/")[1],
				Url:          u.String(),
				LastModified: *object.LastModified,
			}
			out = append(out, media)
		}
	}

	return out, nil
}

func (c *Client) Store(ctx context.Context, filename string, contentBytes []byte) error {
	result, err := c.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(contentBytes),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return err
	}

	if result.Location == "" {
		return errors.New("error store file")
	}

	return nil
}
