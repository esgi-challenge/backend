package usecase

import (
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/email"
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

func (u *userUseCase) GetById(id uint) (*models.User, error) {
	return u.userRepo.GetById(id)
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

func (u *userUseCase) Update(id uint, updatedUser *models.User) (*models.User, error) {
	// Temporary fix for known issue :
	// https://github.com/go-gorm/gorm/issues/5724
	//////////////////////////////////////
	dbUser, err := u.GetById(id)
	if err != nil {
		return nil, err
	}
	updatedUser.CreatedAt = dbUser.CreatedAt
	///////////////////////////////////////

	if dbUser.Email != updatedUser.Email {
		emailM := email.InitEmailManager(u.cfg.Smtp.Username, u.cfg.Smtp.Password, u.cfg.Smtp.Host)
		err = emailM.SendUpdateEmail([]string{dbUser.Email}, updatedUser.Email)
		if err != nil {
			return nil, err
		}
	}

	updatedUser.ID = id
	return u.userRepo.Update(id, updatedUser)
}
