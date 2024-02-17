package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const urlMedia = "https://%s.s3.amazonaws.com/%s"

func (c *Client) List(ctx context.Context, filename string) ([]*Media, error) {
	objects, err := c.client.ListObjects(ctx, &s3.ListObjectsInput{
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
		if len(objects.Contents) == 0 {
			continue
		}

		object := objects.Contents[i]

		if *object.Size != 0 {
			u, _ := url.Parse(fmt.Sprintf(urlMedia, c.bucketName, *object.Key))
			media := &Media{
				Filename:     *object.Key,
				Url:          u.String(),
				LastModified: *object.LastModified,
			}
			out = append(out, media)
		}
	}

	return out, nil
}

func (c *Client) Store(ctx context.Context, filename string, contentBytes []byte) error {
	result, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(contentBytes),
		ACL:    "public-read",
	})
	if err != nil {
		return err
	}

	if result.Location == "" {
		return errors.New("error store file")
	}

	return nil
}
