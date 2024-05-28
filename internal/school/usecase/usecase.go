package usecase

import (
	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type schoolUseCase struct {
	schoolRepo school.Repository
	cfg        *config.Config
	logger     logger.Logger
}

func NewSchoolUseCase(cfg *config.Config, schoolRepo school.Repository, logger logger.Logger) school.UseCase {
	return &schoolUseCase{cfg: cfg, schoolRepo: schoolRepo, logger: logger}
}

func (u *schoolUseCase) Create(school *models.School) (*models.School, error) {
	return u.schoolRepo.Create(school)
}

func (u *schoolUseCase) GetAll() (*[]models.School, error) {
	return u.schoolRepo.GetAll()
}

func (u *schoolUseCase) GetById(id uint) (*models.School, error) {
	return u.schoolRepo.GetById(id)
}

func (u *schoolUseCase) Delete(id uint) error {
	// Check not needed but added to handle a not found error because gorm do not return
	// error if delete on a row that does not exist
	_, err := u.GetById(id)
	if err != nil {
		return err
	}

	return u.schoolRepo.Delete(id)
}
