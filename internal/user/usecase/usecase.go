package usecase

import (
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/google/uuid"
)

type userUseCase struct {
	userRepo user.Repository
	cfg      *config.Config
	logger   logger.Logger
}

func NewUserUseCase(userRepo user.Repository, cfg *config.Config, logger logger.Logger) user.UseCase {
	return &userUseCase{userRepo: userRepo, cfg: cfg, logger: logger}
}

func (u *userUseCase) Create(user *models.User) (*models.User, error) {
	return u.userRepo.Create(user)
}

func (u *userUseCase) GetAll() (*[]models.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUseCase) SendResetMail(email string) (string, error) {
	user, err := u.userRepo.GetByEmail(email)

	if err != nil {
		return "", err
	}

	resetCode := uuid.NewString()

	user.PasswordResetCode = resetCode

	_, err = u.userRepo.Update(
		user.ID,
		user,
	)

	return resetCode, err
}
