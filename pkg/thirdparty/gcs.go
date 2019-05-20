package thirdparty

import (
	"context"

	"cloud.google.com/go/storage"
)

func NewGCSBucket(name string) (*storage.BucketHandle, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	bkt := client.Bucket(name)
	return bkt, nil
}
