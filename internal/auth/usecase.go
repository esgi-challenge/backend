package auth

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Login(payload *models.AuthLogin) (*models.Auth, error)
	Register(user *models.User) (*models.Auth, error)
	InvitationCode(payload *models.AuthInvitationCode) (*models.Auth, error)
	ResetPassword(payload *models.AuthResetPassword) error
}
