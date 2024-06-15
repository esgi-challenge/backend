package usecase

import (
	"errors"
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type schoolUseCase struct {
	schoolRepo school.Repository
	userRepo   user.Repository
	cfg        *config.Config
	logger     logger.Logger
}

func NewSchoolUseCase(cfg *config.Config, schoolRepo school.Repository, userRepo user.Repository, logger logger.Logger) school.UseCase {
	return &schoolUseCase{cfg: cfg, schoolRepo: schoolRepo, userRepo: userRepo, logger: logger}
}

func (u *schoolUseCase) Create(user *models.User, school *models.SchoolCreate) (*models.School, error) {
	existingSchool, err := u.schoolRepo.GetByUser(user)

	if existingSchool != nil {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "You already have a school associated to your account",
		}
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return u.schoolRepo.Create(&models.School{
		UserID: user.ID,
		Name:   school.Name,
	})
}

func (u *schoolUseCase) Invite(user *models.User, schoolInvite *models.SchoolInvite) (*models.User, error) {
	school, err := u.schoolRepo.GetById(schoolInvite.SchoolId)

	if err != nil {
		return nil, err
	}

	if school.UserID != user.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusUnauthorized,
			HttpError:  "Not allowed to perform action on this ressource",
		}
	}

	return u.userRepo.Create(&models.User{
		Email:          schoolInvite.Email,
		Lastname:       schoolInvite.Lastname,
		Firstname:      schoolInvite.Firstname,
		InvitationCode: uuid.NewString(),
		SchoolId:       &school.ID,
	})
}

func (u *schoolUseCase) GetAll() (*[]models.School, error) {
	return u.schoolRepo.GetAll()
}

func (u *schoolUseCase) GetById(id uint) (*models.School, error) {
	return u.schoolRepo.GetById(id)
}

func (u *schoolUseCase) Delete(user *models.User, id uint) error {
	// Check not needed but added to handle a not found error because gorm do not return
	// error if delete on a row that does not exist
	school, err := u.GetById(id)
	if err != nil {
		return err
	}

	if school.UserID != user.ID {
		return errorHandler.HttpError{
			HttpStatus: http.StatusUnauthorized,
			HttpError:  "Not allowed to perform action on this ressource",
		}
	}

	return u.schoolRepo.Delete(id)
}
