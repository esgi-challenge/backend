package project

import (
	"github.com/esgi-challenge/backend/internal/models"
)

type Repository interface {
	Create(project *models.Project) (*models.Project, error)
	GetAll(user *models.User) (*[]models.Project, error)
  GetAllByTeacher(user *models.User) (*[]models.Project, error)
	GetById(user *models.User, id uint) (*models.Project, error)
	GetGroups(user *models.User, id uint) (*[]models.ProjectStudent, error)
	JoinProject(project *models.ProjectStudent) (*models.ProjectStudent, error)
	GetJoined(projectId uint, userId uint) (*[]models.ProjectStudent, error)
	DeleteJoined(projectId uint, userId uint) error
	Update(id uint, project *models.Project) (*models.Project, error)
	Delete(id uint) error
}
