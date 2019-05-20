package manager

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/Bo0km4n/polybucket/pkg/thirdparty"
)

const (
	historyFileName = "polybucket_hisotory.json"
)

type GCS struct {
	history *history
	bucket  *storage.BucketHandle
}

func NewGCSManager(bucket, prefix string) (*GCS, error) {
	bkt, err := thirdparty.NewGCSBucket(bucket)
	if err != nil {
		return nil, err
	}
	historyFileName := strings.TrimRight(prefix, "/") + "/" + historyFileName

	history, err := fetchHistory(bkt, historyFileName)
	if err != nil {
		return nil, err
	}
	return &GCS{
		history: history,
		bucket:  bkt,
	}, nil
}

func fetchHistory(bucket *storage.BucketHandle, fileName string) (*history, error) {
	obj := bucket.Object(fileName)
	ctx := context.Background()
	if _, err := obj.Attrs(context.Background()); err != nil {
		// create history file
		writer := obj.NewWriter(ctx)
		defer writer.Close()
		newHistory := &history{
			Versions: make(map[string]string),
		}
		b, err := json.Marshal(newHistory)
		if err != nil {
			return nil, err
		}
		if _, err := writer.Write(b); err != nil {
			return nil, err
		}
		return newHistory, nil
	} else {
		// load exitst history
		reader, err := obj.NewReader(ctx)
		if err != nil {
			return &history{}, err
		}
		defer reader.Close()
		b, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		latestHistory := &history{}
		if err := json.Unmarshal(b, latestHistory); err != nil {
			return nil, err
		}
		return latestHistory, nil
	}
}
