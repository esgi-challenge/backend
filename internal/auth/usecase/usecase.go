package usecase

import (
	"errors"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/auth"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/jwt"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type authUseCase struct {
	userRepo user.Repository
	cfg      *config.Config
	logger   logger.Logger
}

func NewAuthUseCase(cfg *config.Config, userRepo user.Repository, logger logger.Logger) auth.UseCase {
	return &authUseCase{cfg: cfg, userRepo: userRepo, logger: logger}
}

func (u *authUseCase) Login(payload *models.AuthLogin) (*models.Auth, error) {
	user, err := u.userRepo.GetByEmail(payload.Email)

	if err != nil {
		return nil, errors.New("authentication unaccepted")
	}

	token, err := jwt.Generate(u.cfg.JwtSecret, user)

	if err != nil {
		return nil, err
	}

	return &models.Auth{
		Token: token,
	}, nil
}

func (u *authUseCase) Register(payload *models.AuthRegister) (*models.Auth, error) {

	user, err := u.userRepo.Create(&models.User{
		Lastname:  payload.Lastname,
		Firstname: payload.Firstname,
		Email:     payload.Email,
		Password:  payload.Password,
		UserKind:  models.ADMINISTRATOR,
	})

	if err != nil {
		return nil, err
	}

	token, err := jwt.Generate(u.cfg.JwtSecret, user)

	if err != nil {
		return nil, err
	}

	return &models.Auth{
		Token: token,
	}, nil
}
