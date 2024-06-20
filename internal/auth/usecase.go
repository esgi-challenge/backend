package auth

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Login(payload *models.AuthLogin) (*models.Auth, error)
	Register(payload *models.AuthRegister) (*models.Auth, error)
	ResetPassword(payload *models.AuthResetPassword) error
}
