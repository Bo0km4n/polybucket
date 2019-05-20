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
	historyFileName = "_polybucket/hisotory.json"
)

type GCS struct {
	history *history
	bucket  *storage.BucketHandle
	prefix  string
	objects []*storage.ObjectHandle
}

func NewGCSManager(bucket, prefix, modelName string) (*GCS, error) {
	prefix = strings.TrimRight(prefix, "/") + "/"
	bkt, err := thirdparty.NewGCSBucket(bucket)
	if err != nil {
		return nil, err
	}
	historyFileName := prefix + "/" + historyFileName

	history, err := fetchHistory(bkt, historyFileName)
	if err != nil {
		return nil, err
	}

	return &GCS{
		history: history,
		bucket:  bkt,
		prefix:  prefix,
	}, nil
}

func (gcs *GCS) LazyFetchModel(version string) (*storage.ObjectHandle, error) {
	// generation, ok := gcs.history.Versions[version]
	// if !ok {
	// 	return nil, fmt.Errorf("Not found version: %s", version)
	// }
	// objs := gcs.bucket.Objects(context.Background(), &storage.Query{
	// 	Prefix:    gcs.prefix,
	// 	Delimiter: "/",
	// 	Versions:  true,
	// })
	// for {
	// 	obj, err := objs.Next()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	return nil, nil
}

func fetchHistory(bucket *storage.BucketHandle, fileName string) (*history, error) {
	obj := bucket.Object(fileName)
	ctx := context.Background()
	if _, err := obj.Attrs(context.Background()); err != nil {
		// create history file
		writer := obj.NewWriter(ctx)
		defer writer.Close()
		newHistory := &history{
			Versions: make(map[string]int64),
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
