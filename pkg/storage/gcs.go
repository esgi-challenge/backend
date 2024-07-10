package storage

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Storage struct {
	cfg    *config.Config
	psqlDB *gorm.DB
	logger logger.Logger
	client *storage.Client
}

func NewStorage(cfg *config.Config, psqlDB *gorm.DB, logger logger.Logger) *Storage {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	logger.Info("Storage: Setting up...")
	if err != nil {
		logger.Error("Storage: ")
		return nil
	}
	logger.Info("Storage: Set up!")

	return &Storage{
		cfg:    cfg,
		psqlDB: psqlDB,
		logger: logger,
		client: client,
	}
}

func (s Storage) UploadFile(ctx context.Context, file []byte) (string, error) {
	bkt := s.client.Bucket(s.cfg.Bucket)

	filename := fmt.Sprintf("files/%s", uuid.NewString())

	obj := bkt.Object(filename)

	w := obj.NewWriter(ctx)
	w.ACL = []storage.ACLRule{{
		Entity: storage.AllUsers,
		Role:   storage.RoleReader,
	}}
	_, err := w.Write(file)

	if err != nil {
		s.logger.Error("GCS Error: %s", err)
		return "", err
	}

	if err := w.Close(); err != nil {
		s.logger.Error("GCS Error: %s", err)
		return "", err
	}

	return filename, nil
}
