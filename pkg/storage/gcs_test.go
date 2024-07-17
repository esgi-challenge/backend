package storage

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/fsouza/fake-gcs-server/fakestorage"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"google.golang.org/api/option"
)

func setup() (*Storage, *fakestorage.Server, error) {
	cfg := &config.Config{
		Bucket: "test-bucket",
	}

	log := logger.NewLogger()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	server := fakestorage.NewServer([]fakestorage.Object{})

	serverURL, err := url.Parse(server.URL())
	if err != nil {
		return nil, nil, err
	}

	client, err := storage.NewClient(context.Background(), option.WithHTTPClient(&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(serverURL),
		},
	}))
	if err != nil {
		return nil, nil, err
	}

	storage := &Storage{
		cfg:    cfg,
		psqlDB: db,
		logger: log,
		client: client,
	}

	return storage, server, nil
}

func TestNewStorage(t *testing.T) {
	storage, server, err := setup()
	defer server.Stop()

	assert.NoError(t, err)
	assert.NotNil(t, storage)
	assert.NotNil(t, storage.client)
}
