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

func (u *schoolUseCase) RemoveUser(userId uint, userKind models.UserKind, school *models.School) error {
	var users *[]models.User
	var err error
	if userKind == models.STUDENT {
		users, err = u.schoolRepo.GetSchoolStudents(school.ID)
	} else if userKind == models.TEACHER {
		users, err = u.schoolRepo.GetSchoolTeachers(school.ID)
	} else {
		return gorm.ErrRecordNotFound
	}
	if err != nil {
		return err
	}

	isInSchool := false

	for _, v := range *users {
		if v.ID == userId {
			isInSchool = true
		}
	}

	if !isInSchool {
		return gorm.ErrRecordNotFound
	}

	err = u.userRepo.Delete(userId)

	return err
}

func (u *schoolUseCase) AddUser(user *models.User) (*models.User, error) {
	return u.userRepo.Create(user)
}

func (u *schoolUseCase) Invite(user *models.User, schoolInvite *models.SchoolInvite) (*models.User, error) {
	var userKind models.UserKind = models.STUDENT

	school, err := u.schoolRepo.GetByUser(user)

	if err != nil {
		return nil, err
	}

	if schoolInvite.Type == "TEACHER" {
		userKind = models.TEACHER
	}

	return u.userRepo.Create(&models.User{
		Email:          schoolInvite.Email,
		Lastname:       schoolInvite.Lastname,
		Firstname:      schoolInvite.Firstname,
		InvitationCode: uuid.NewString(),
		SchoolId:       &school.ID,
		UserKind:       &userKind,
	})
}

func (u *schoolUseCase) GetAll() (*[]models.School, error) {
	return u.schoolRepo.GetAll()
}

func (u *schoolUseCase) GetById(id uint) (*models.School, error) {
	return u.schoolRepo.GetById(id)
}

func (u *schoolUseCase) GetByUser(user *models.User) (*models.School, error) {
	return u.schoolRepo.GetByUser(user)
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

func (u *schoolUseCase) GetSchoolStudents(schoolId uint) (*[]models.User, error) {
	return u.schoolRepo.GetSchoolStudents(schoolId)
}

func (u *schoolUseCase) GetSchoolTeachers(schoolId uint) (*[]models.User, error) {
	return u.schoolRepo.GetSchoolTeachers(schoolId)
}
