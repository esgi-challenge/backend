package usecase

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/path"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type pathUseCase struct {
	pathRepo   path.Repository
	schoolRepo school.Repository
	cfg        *config.Config
	logger     logger.Logger
}

func NewPathUseCase(cfg *config.Config, pathRepo path.Repository, schoolRepo school.Repository, logger logger.Logger) path.UseCase {
	return &pathUseCase{cfg: cfg, pathRepo: pathRepo, schoolRepo: schoolRepo, logger: logger}
}

func (u *pathUseCase) Create(user *models.User, path *models.Path) (*models.Path, error) {
	school, err := u.schoolRepo.GetById(path.SchoolId)

	if err != nil {
		return nil, err
	}

	if school.UserID != user.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This school is not yours",
		}
	}

	return u.pathRepo.Create(path)
}

func (u *pathUseCase) GetAll() (*[]models.Path, error) {
	return u.pathRepo.GetAll()
}

func (u *pathUseCase) GetById(id uint) (*models.Path, error) {
	return u.pathRepo.GetById(id)
}

func (u *pathUseCase) Update(user *models.User, id uint, updatedPath *models.Path) (*models.Path, error) {
	school, err := u.schoolRepo.GetById(updatedPath.SchoolId)

	if err != nil {
		return nil, err
	}

	if school.UserID != user.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This school is not yours",
		}
	}

	// Temporary fix for known issue :
	// https://github.com/go-gorm/gorm/issues/5724
	//////////////////////////////////////
	dbPath, err := u.GetById(id)

	if dbPath.SchoolId != school.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusBadRequest,
			HttpError:  errorHandler.BadRequest.Error(),
		}
	}

	if err != nil {
		return nil, err
	}
	updatedPath.CreatedAt = dbPath.CreatedAt
	///////////////////////////////////////

	updatedPath.ID = id
	return u.pathRepo.Update(id, updatedPath)
}

func (u *pathUseCase) Delete(user *models.User, id uint) error {
	path, err := u.GetById(id)

	if err != nil {
		return err
	}

	school, err := u.schoolRepo.GetById(path.SchoolId)

	if err != nil {
		return err
	}

	if school.UserID != user.ID {
		return errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This school is not yours",
		}
	}

	return u.pathRepo.Delete(id)
}
