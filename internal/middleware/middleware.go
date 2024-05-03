package middleware

import (
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type MiddlewareManager struct {
  cfg *config.Config
  logger logger.Logger
}

func InitMiddlewareManager(cfg *config.Config, logger logger.Logger) *MiddlewareManager {
  return &MiddlewareManager{cfg: cfg, logger: logger}
}
