package usecase

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/informations"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type informationsUseCase struct {
	informationsRepo informations.Repository
	schoolRepo       school.Repository
	cfg              *config.Config
	logger           logger.Logger
}

func NewInformationsUseCase(cfg *config.Config, informationsRepo informations.Repository, schoolRepo school.Repository, logger logger.Logger) informations.UseCase {
	return &informationsUseCase{
		cfg:              cfg,
		informationsRepo: informationsRepo,
		schoolRepo:       schoolRepo,
		logger:           logger,
	}
}

func (u *informationsUseCase) Create(user *models.User, informations *models.Informations) (*models.Informations, error) {
	school, err := u.schoolRepo.GetById(informations.SchoolId)

	if err != nil {
		return nil, err
	}

	if school.UserID != user.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This school is not yours",
		}
	}

	return u.informationsRepo.Create(informations)
}

func (u *informationsUseCase) GetAll(user *models.User) (*[]models.Informations, error) {
	if *user.UserKind == models.STUDENT {
		return u.informationsRepo.GetBySchoolId(*user.SchoolId)
	}

  school, err := u.schoolRepo.GetByUser(user)
  if err != nil {
    return nil, err
  }

	return u.informationsRepo.GetBySchoolId(school.ID)
}

func (u *informationsUseCase) GetById(user *models.User, id uint) (*models.Informations, error) {
	if *user.UserKind == models.STUDENT {
		return u.informationsRepo.GetBySchoolIdAndId(*user.SchoolId, id)
	}

	return u.informationsRepo.GetById(id)
}

func (u *informationsUseCase) Delete(user *models.User, id uint) error {
	// Check not needed but added to handle a not found error because gorm do not return
	// error if delete on a row that does not exist
	dbInformation, err := u.GetById(user, id)

	if err != nil {
		return err
	}

	school, err := u.schoolRepo.GetById(dbInformation.SchoolId)

	if err != nil {
		return err
	}

	if school.UserID != user.ID {
		return errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This information is not yours",
		}
	}

	return u.informationsRepo.Delete(id)
}
