package project

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type UseCase interface {
	Create(user *models.User, project *models.Project) (*models.Project, error)
	GetAll(user *models.User) (*[]models.Project, error)
	GetById(user *models.User, id uint) (*models.Project, error)
	GetGroups(user *models.User, id uint) ([]*models.ProjectGroup, error)
	JoinProject(user *models.User, join *models.ProjectStudentCreate, id uint) (*models.ProjectStudent, error)
	QuitProject(user *models.User, id uint) error
	Update(user *models.User, id uint, updatedProject *models.Project) (*models.Project, error)
	Delete(user *models.User, id uint) error
}
