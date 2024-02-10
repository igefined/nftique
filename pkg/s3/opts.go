package s3

import "time"

type Media struct {
	Filename     string
	Date         string
	Url          string
	LastModified time.Time
}

type Opt func(client *Client)

func WithBucketName(bucketName string) Opt {
	return func(client *Client) {
		client.bucketName = bucketName
	}
}
