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

	cfg    *config.Config
	logger logger.Logger
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

func (u *pathUseCase) GetAllBySchoolId(schoolId uint) (*[]models.Path, error) {
	return u.pathRepo.GetAllBySchoolId(schoolId)
}

func (u *pathUseCase) GetById(id uint) (*models.Path, error) {
	return u.pathRepo.GetById(id)
}

func (u *pathUseCase) Update(id uint, updatedPath *models.Path) (*models.Path, error) {
	// Temporary fix for known issue :
	// https://github.com/go-gorm/gorm/issues/5724
	//////////////////////////////////////
	dbPath, err := u.GetById(id)

	if err != nil {
		return nil, err
	}
	updatedPath.CreatedAt = dbPath.CreatedAt
	///////////////////////////////////////

	updatedPath.ID = id
	return u.pathRepo.Update(id, updatedPath)
}

func (u *pathUseCase) Delete(id uint) error {
	return u.pathRepo.Delete(id)
}
