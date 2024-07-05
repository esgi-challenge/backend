package usecase

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/campus"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type campusUseCase struct {
	campusRepo campus.Repository
	schoolRepo school.Repository
	cfg        *config.Config
	logger     logger.Logger
}

func NewCampusUseCase(cfg *config.Config, campusRepo campus.Repository, schoolRepo school.Repository, logger logger.Logger) campus.UseCase {
	return &campusUseCase{cfg: cfg, campusRepo: campusRepo, schoolRepo: schoolRepo, logger: logger}
}

func (u *campusUseCase) Create(user *models.User, campus *models.Campus) (*models.Campus, error) {
	school, err := u.schoolRepo.GetById(campus.SchoolId)

	if err != nil {
		return nil, err
	}

	if school.UserID != user.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This school is not yours",
		}
	}

	return u.campusRepo.Create(campus)
}

func (u *campusUseCase) GetAll() (*[]models.Campus, error) {
	return u.campusRepo.GetAll()
}

func (u *campusUseCase) GetById(id uint) (*models.Campus, error) {
	return u.campusRepo.GetById(id)
}

func (u *campusUseCase) GetAllBySchoolId(schoolId uint) (*[]models.Campus, error) {
	return u.campusRepo.GetAllBySchoolId(schoolId)
}

func (u *campusUseCase) Update(user *models.User, id uint, updatedCampus *models.Campus) (*models.Campus, error) {
	school, err := u.schoolRepo.GetById(updatedCampus.SchoolId)

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
	dbCampus, err := u.GetById(id)
	if err != nil {
		return nil, err
	}
	updatedCampus.CreatedAt = dbCampus.CreatedAt
	///////////////////////////////////////

	updatedCampus.ID = id
	return u.campusRepo.Update(id, updatedCampus)
}

func (u *campusUseCase) Delete(user *models.User, id uint) error {
	campus, err := u.GetById(id)

	if err != nil {
		return err
	}

	school, err := u.schoolRepo.GetById(campus.SchoolId)

	if err != nil {
		return err
	}

	if school.UserID != user.ID {
		return errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "This school is not yours",
		}
	}

	return u.campusRepo.Delete(id)
}
