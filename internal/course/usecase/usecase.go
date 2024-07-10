package usecase

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/course"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/path"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type courseUseCase struct {
	courseRepo course.Repository
	pathRepo   path.Repository
	schoolRepo school.Repository
	cfg        *config.Config
	logger     logger.Logger
}

func NewCourseUseCase(cfg *config.Config, courseRepo course.Repository, pathRepo path.Repository, schoolRepo school.Repository, logger logger.Logger) course.UseCase {
	return &courseUseCase{
		cfg:        cfg,
		courseRepo: courseRepo,
		pathRepo:   pathRepo,
		schoolRepo: schoolRepo,
		logger:     logger,
	}
}

func (u *courseUseCase) Create(user *models.User, course *models.Course) (*models.Course, error) {
	path, err := u.pathRepo.GetById(course.PathId)

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
			HttpError:  "This path is not yours",
		}
	}

	return u.courseRepo.Create(course)
}

func (u *courseUseCase) GetAll() (*[]models.Course, error) {
	return u.courseRepo.GetAll()
}

func (u *courseUseCase) GetAllBySchoolId(schoolId uint) (*[]models.Course, error) {
	return u.courseRepo.GetAllBySchoolId(schoolId)
}

func (u *courseUseCase) GetById(id uint) (*models.Course, error) {
	return u.courseRepo.GetById(id)
}

func (u *courseUseCase) Update(id uint, updatedCourse *models.Course) (*models.Course, error) {
	// Temporary fix for known issue :
	// https://github.com/go-gorm/gorm/issues/5724
	//////////////////////////////////////
	dbCourse, err := u.GetById(id)
	if err != nil {
		return nil, err
	}
	updatedCourse.CreatedAt = dbCourse.CreatedAt
	///////////////////////////////////////

	updatedCourse.ID = id
	return u.courseRepo.Update(id, updatedCourse)
}

func (u *courseUseCase) Delete(id uint) error {
	return u.courseRepo.Delete(id)
}
