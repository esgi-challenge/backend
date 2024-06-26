package usecase

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/course"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/path"
	"github.com/esgi-challenge/backend/internal/schedule"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type scheduleUseCase struct {
	scheduleRepo schedule.Repository
	courseRepo   course.Repository
	schoolRepo   school.Repository
	pathRepo     path.Repository
	cfg          *config.Config
	logger       logger.Logger
}

func NewScheduleUseCase(cfg *config.Config, scheduleRepo schedule.Repository, courseRepo course.Repository, pathRepo path.Repository, schoolRepo school.Repository, logger logger.Logger) schedule.UseCase {
	return &scheduleUseCase{
		cfg:          cfg,
		scheduleRepo: scheduleRepo,
		courseRepo:   courseRepo,
		schoolRepo:   schoolRepo,
		pathRepo:     pathRepo,
		logger:       logger,
	}
}

func (u *scheduleUseCase) Create(user *models.User, schedule *models.Schedule) (*models.Schedule, error) {
	course, err := u.courseRepo.GetById(schedule.CourseId)

	if err != nil {
		return nil, err
	}

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
			HttpError:  "This course is not yours",
		}
	}

	return u.scheduleRepo.Create(schedule)
}

func (u *scheduleUseCase) GetAll(user *models.User) (*[]models.Schedule, error) {
	return u.scheduleRepo.GetAll()
}

func (u *scheduleUseCase) GetById(user *models.User, id uint) (*models.Schedule, error) {
	return u.scheduleRepo.GetById(id)
}

func (u *scheduleUseCase) Update(user *models.User, id uint, updatedSchedule *models.Schedule) (*models.Schedule, error) {
	// Temporary fix for known issue :
	// https://github.com/go-gorm/gorm/issues/5724
	//////////////////////////////////////
	dbSchedule, err := u.GetById(user, id)
	if err != nil {
		return nil, err
	}

	updatedSchedule.CreatedAt = dbSchedule.CreatedAt
	///////////////////////////////////////
	course, err := u.courseRepo.GetById(dbSchedule.CourseId)

	if err != nil {
		return nil, err
	}

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
			HttpError:  "This course is not yours",
		}
	}

	course, err = u.courseRepo.GetById(updatedSchedule.CourseId)

	if err != nil {
		return nil, err
	}

	path, err = u.pathRepo.GetById(course.PathId)

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
			HttpError:  "This course is not yours",
		}
	}

	updatedSchedule.ID = id
	return u.scheduleRepo.Update(id, updatedSchedule)
}

func (u *scheduleUseCase) Delete(user *models.User, id uint) error {
	// Check not needed but added to handle a not found error because gorm do not return
	// error if delete on a row that does not exist
	schedule, err := u.GetById(user, id)

	if err != nil {
		return err
	}

	course, err := u.courseRepo.GetById(schedule.CourseId)

	if err != nil {
		return err
	}

	path, err := u.pathRepo.GetById(course.PathId)

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
			HttpError:  "This course is not yours",
		}
	}

	return u.scheduleRepo.Delete(id)
}
