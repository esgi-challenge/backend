package usecase

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/class"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/path"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type classUseCase struct {
	classRepo  class.Repository
	schoolRepo school.Repository
	pathRepo   path.Repository
	userRepo   user.Repository
	cfg        *config.Config
	logger     logger.Logger
}

func NewClassUseCase(cfg *config.Config, classRepo class.Repository, pathRepo path.Repository, schoolRepo school.Repository, userRepo user.Repository, logger logger.Logger) class.UseCase {
	return &classUseCase{
		cfg:        cfg,
		classRepo:  classRepo,
		pathRepo:   pathRepo,
		schoolRepo: schoolRepo,
		userRepo:   userRepo,
		logger:     logger,
	}
}

func (u *classUseCase) Create(user *models.User, class *models.Class) (*models.Class, error) {
	path, err := u.pathRepo.GetById(class.PathId)

	if err != nil {
		return nil, err
	}

	school, err := u.schoolRepo.GetById(path.SchoolId)

	if err != nil {
		return nil, err
	}

	if school.UserID != user.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  errorHandler.Forbidden.Error(),
		}
	}

	return u.classRepo.Create(class)
}

func (u *classUseCase) GetAll() (*[]models.Class, error) {
	return u.classRepo.GetAll()
}

func (u *classUseCase) GetById(id uint) (*models.Class, error) {
	return u.classRepo.GetById(id)
}

func (u *classUseCase) Add(user *models.User, id uint, addClass *models.ClassAdd) (*models.Class, error) {
	student, err := u.userRepo.GetById(*addClass.UserId)

	if err != nil {
		return nil, err
	}

	class, err := u.GetById(id)

	if err != nil {
		return nil, err
	}

	contains := false

	for _, k := range class.Students {
		if k.ID == student.ID {
			contains = true
			break
		}
	}

	if contains {
		class.Students = append(class.Students, *student)
	}

	return u.Update(user, id, class)
}

func (u *classUseCase) Update(user *models.User, id uint, updatedClass *models.Class) (*models.Class, error) {
	path, err := u.pathRepo.GetById(updatedClass.PathId)

	if err != nil {
		return nil, err
	}

	school, err := u.schoolRepo.GetById(path.SchoolId)

	if err != nil {
		return nil, err
	}

	if school.UserID != user.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  errorHandler.Forbidden.Error(),
		}
	}

	// Temporary fix for known issue :
	// https://github.com/go-gorm/gorm/issues/5724
	//////////////////////////////////////
	dbClass, err := u.GetById(id)

	if err != nil {
		return nil, err
	}

	path, err = u.pathRepo.GetById(dbClass.PathId)

	if err != nil {
		return nil, err
	}

	school, err = u.schoolRepo.GetById(path.SchoolId)

	if err != nil {
		return nil, err
	}

	if school.UserID != user.ID {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  errorHandler.Forbidden.Error(),
		}
	}

	updatedClass.CreatedAt = dbClass.CreatedAt
	///////////////////////////////////////

	updatedClass.ID = id
	return u.classRepo.Update(id, updatedClass)
}

func (u *classUseCase) Delete(user *models.User, id uint) error {
	// Check not needed but added to handle a not found error because gorm do not return
	// error if delete on a row that does not exist
	class, err := u.GetById(id)

	if err != nil {
		return err
	}

	path, err := u.pathRepo.GetById(class.PathId)

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
			HttpError:  errorHandler.Forbidden.Error(),
		}
	}

	if err != nil {
		return err
	}

	return u.classRepo.Delete(id)
}
