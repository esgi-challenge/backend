package usecase

import (
	"net/http"
	"time"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/campus"
	"github.com/esgi-challenge/backend/internal/course"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/path"
	"github.com/esgi-challenge/backend/internal/schedule"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/google/uuid"
)

type scheduleUseCase struct {
	scheduleRepo schedule.Repository
	courseRepo   course.Repository
	schoolRepo   school.Repository
	pathRepo     path.Repository
	campusRepo   campus.Repository
	userRepo     user.Repository
	cfg          *config.Config
	logger       logger.Logger
}

func NewScheduleUseCase(cfg *config.Config, scheduleRepo schedule.Repository, courseRepo course.Repository, pathRepo path.Repository, schoolRepo school.Repository, campusRepo campus.Repository, userRepo user.Repository, logger logger.Logger) schedule.UseCase {
	return &scheduleUseCase{
		cfg:          cfg,
		scheduleRepo: scheduleRepo,
		courseRepo:   courseRepo,
		schoolRepo:   schoolRepo,
		campusRepo:   campusRepo,
		userRepo:     userRepo,
		pathRepo:     pathRepo,
		logger:       logger,
	}
}

func (u *scheduleUseCase) Create(user *models.User, schedule *models.ScheduleCreate) (*models.Schedule, error) {
	course, err := u.courseRepo.GetById(*schedule.CourseId)

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

	if int64(*schedule.Time) < time.Now().Unix() {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusBadRequest,
			HttpError:  "You cannot create a schedule in the past",
		}
	}

	if *schedule.Duration <= 0 {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusBadRequest,
			HttpError:  "You cannot create a negative schedule duration",
		}
	}

	return u.scheduleRepo.Create(&models.Schedule{
		Time:          *schedule.Time,
		Duration:      *schedule.Duration,
		SignatureCode: uuid.NewString(),
		QrCodeEnabled: *schedule.QrCodeEnabled,
		CourseId:      *schedule.CourseId,
		CampusId:      *schedule.CampusId,
		ClassId:       *schedule.ClassId,
		SchoolId:      school.ID,
	})
}

func (u *scheduleUseCase) Sign(signature *models.ScheduleSignatureCreate, user *models.User, scheduleId uint) (*models.ScheduleSignature, error) {
	var kind models.SignatureKind

	if *user.UserKind != 0 {
		kind = models.SIGNATURE_ADMINISTRATOR
	} else {
		kind = models.SIGNATURE_STUDENT
	}

	schedule, err := u.scheduleRepo.GetById(user, scheduleId)

	if err != nil {
		return nil, err
	}

	if schedule.SignatureCode != signature.SignatureCode {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusBadRequest,
			HttpError:  "The signature code is not correct",
		}
	}

	signingStudent := *user

	if *user.UserKind != 0 {
		student, err := u.userRepo.GetById(signature.UserId)
		if err != nil {
			return nil, err
		}

		signingStudent = *student
	}

	return u.scheduleRepo.Sign(&models.ScheduleSignature{
		Student:  signingStudent,
		Schedule: *schedule,
		Kind:     kind,
	})
}

func (u *scheduleUseCase) GetAll(user *models.User) (*[]models.ScheduleGet, error) {
	schedules, err := u.GetAllByUser(user)

	if err != nil {
		return nil, err
	}

	var finalSchedules []models.ScheduleGet

	for _, schedule := range *schedules {
		schedulePreload, err := u.GetPreloadById(schedule.ID)
		if err != nil {
			return nil, err
		}

		finalSchedules = append(finalSchedules, models.ScheduleGet{
			Schedule: schedule,
			Campus:   schedulePreload.Campus,
			Course:   schedulePreload.Course,
		})

	}

	return &finalSchedules, nil
}

func (u *scheduleUseCase) GetUnattended(user *models.User) (*[]models.ScheduleGet, error) {
	now := time.Now()
	timestamp := uint(now.Unix())
	unattendedSchedules := []models.ScheduleGet{}
	schedules, err := u.GetAll(user)

	if err != nil {
		return nil, err
	}

	for _, schedule := range *schedules {
		if schedule.Schedule.Time+(schedule.Schedule.Duration*60) < timestamp {
			_, err := u.CheckSign(user, schedule.Schedule.ID)

			if err != nil {
				unattendedSchedules = append(unattendedSchedules, schedule)
			}

		}
	}

	return &unattendedSchedules, nil
}

func (u *scheduleUseCase) CheckSign(user *models.User, scheduleId uint) (*models.ScheduleSignature, error) {
	return u.scheduleRepo.GetSign(user.ID, scheduleId)
}

func (u *scheduleUseCase) GetStudentsSignature(user *models.User, scheduleId uint) (*models.ScheduleSignatureGet, error) {
	schedule, err := u.scheduleRepo.GetById(user, scheduleId)

	if err != nil {
		return nil, err
	}

	students, err := u.scheduleRepo.GetScheduleStudents(schedule.ClassId)

	if err != nil {
		return nil, err
	}

	signatures, err := u.scheduleRepo.GetScheduleSignatures(schedule.ID)

	if err != nil {
		return nil, err
	}

	return &models.ScheduleSignatureGet{
		Students:  *students,
		Signature: *signatures,
	}, nil
}

func (u *scheduleUseCase) GetSignatureCode(user *models.User, scheduleId uint) (*models.ScheduleSignatureCode, error) {
	schedule, err := u.scheduleRepo.GetById(user, scheduleId)

	if err != nil {
		return nil, err
	}

	return &models.ScheduleSignatureCode{
		SignatureCode: schedule.SignatureCode,
	}, nil
}

func (u *scheduleUseCase) GetAllByUser(user *models.User) (*[]models.Schedule, error) {
	if uint(*user.UserKind) == models.ADMINISTRATOR {
		school, err := u.schoolRepo.GetByUser(user)

		if err != nil {
			return nil, err
		}

		return u.scheduleRepo.GetAllBySchoolId(school.ID)
	} else if uint(*user.UserKind) == models.TEACHER {
		return u.scheduleRepo.GetAllByTeacherId(user.ID)

	} else {
		return u.scheduleRepo.GetAllByClassId(*user.ClassRefer)

	}
}

func (u *scheduleUseCase) GetPreloadById(scheduleId uint) (*models.Schedule, error) {
	return u.scheduleRepo.GetPreloadById(scheduleId)
}

func (u *scheduleUseCase) GetById(user *models.User, id uint) (*models.ScheduleGet, error) {
	schedule, err := u.scheduleRepo.GetById(user, id)

	if err != nil {
		return nil, err
	}

	scheduleWithPreload, err := u.GetPreloadById(schedule.ID)
	if err != nil {
		return nil, err
	}

	return &models.ScheduleGet{
		Schedule: *schedule,
		Campus:   scheduleWithPreload.Campus,
		Course:   scheduleWithPreload.Course,
	}, nil
}

func (u *scheduleUseCase) Update(user *models.User, id uint, updatedSchedule *models.Schedule) (*models.Schedule, error) {

	if int64(updatedSchedule.Time) < time.Now().Unix() {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusBadRequest,
			HttpError:  "You cannot create a schedule in the past",
		}
	}

	if updatedSchedule.Duration <= 0 {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusBadRequest,
			HttpError:  "You cannot create a negative schedule duration",
		}
	}

	// Temporary fix for known issue :
	// https://github.com/go-gorm/gorm/issues/5724
	//////////////////////////////////////
	dbSchedule, err := u.GetById(user, id)
	if err != nil {
		return nil, err
	}

	updatedSchedule.CreatedAt = dbSchedule.Schedule.CreatedAt
	///////////////////////////////////////
	school, err := u.schoolRepo.GetById(updatedSchedule.SchoolId)

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
	_, err := u.GetById(user, id)

	if err != nil {
		return err
	}

	return u.scheduleRepo.Delete(id)
}
