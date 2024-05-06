package usecase

import (
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type userUseCase struct {
	userRepo user.Repository
	cfg      *config.Config
	logger   logger.Logger
}

func NewUserUseCase(userRepo user.Repository, cfg *config.Config, logger logger.Logger) user.UseCase {
	return &userUseCase{userRepo: userRepo, cfg: cfg, logger: logger}
}

func (u *userUseCase) Create(ctx *gin.Context, user *models.User) (*models.User, error) {
	return u.userRepo.Create(user)
}
