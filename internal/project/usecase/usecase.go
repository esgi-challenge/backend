package usecase

import (
	"net/http"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/class"
	"github.com/esgi-challenge/backend/internal/course"
	"github.com/esgi-challenge/backend/internal/document"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/project"
	"github.com/esgi-challenge/backend/pkg/errorHandler"
	"github.com/esgi-challenge/backend/pkg/logger"
)

type projectUseCase struct {
	projectRepo     project.Repository
	courseUseCase   course.UseCase
	classUseCase    class.UseCase
	documentUseCase document.UseCase
	cfg             *config.Config
	logger          logger.Logger
}

func NewProjectUseCase(cfg *config.Config, projectRepo project.Repository, courseUseCase course.UseCase, classUseCase class.UseCase, documentUseCase document.UseCase, logger logger.Logger) project.UseCase {
	return &projectUseCase{
		cfg:             cfg,
		projectRepo:     projectRepo,
		courseUseCase:   courseUseCase,
		classUseCase:    classUseCase,
		documentUseCase: documentUseCase,
		logger:          logger,
	}
}

func (u *projectUseCase) Create(user *models.User, project *models.Project) (*models.Project, error) {
	_, err := u.courseUseCase.GetById(project.CourseId)

	if err != nil {
		return nil, err
	}

	_, err = u.documentUseCase.GetById(user, project.DocumentId)

	if err != nil {
		return nil, err
	}

	_, err = u.classUseCase.GetById(project.ClassId)

	if err != nil {
		return nil, err
	}

	return u.projectRepo.Create(project)
}

func (u *projectUseCase) JoinProject(user *models.User, join *models.ProjectStudentCreate, id uint) (*models.ProjectStudent, error) {
	project, err := u.GetById(user, id)

	if err != nil {
		return nil, err
	}

	joined, err := u.projectRepo.GetJoined(project.ID, user.ID)

	if err != nil {
		return nil, err
	}

	if len(*joined) != 0 {
		return nil, errorHandler.HttpError{
			HttpStatus: http.StatusForbidden,
			HttpError:  "You already joined a group",
		}
	}

	return u.projectRepo.JoinProject(&models.ProjectStudent{
		Group:     *join.Group,
		ProjectId: project.ID,
		StudentId: user.ID,
	})
}

func (u *projectUseCase) QuitProject(user *models.User, id uint) error {
	project, err := u.GetById(user, id)

	if err != nil {
		return err
	}

	err = u.projectRepo.DeleteJoined(project.ID, user.ID)

	return err
}

func (u *projectUseCase) GetAll(user *models.User) (*[]models.Project, error) {
	return u.projectRepo.GetAll(user)
}

func (u *projectUseCase) GetById(user *models.User, id uint) (*models.Project, error) {
	return u.projectRepo.GetById(user, id)
}

func (u *projectUseCase) Update(user *models.User, id uint, updatedProject *models.Project) (*models.Project, error) {
	_, err := u.courseUseCase.GetById(updatedProject.CourseId)

	if err != nil {
		return nil, err
	}

	_, err = u.classUseCase.GetById(updatedProject.ClassId)

	if err != nil {
		return nil, err
	}

	// Temporary fix for known issue :
	// https://github.com/go-gorm/gorm/issues/5724
	//////////////////////////////////////
	dbProject, err := u.GetById(user, id)
	if err != nil {
		return nil, err
	}

	updatedProject.CreatedAt = dbProject.CreatedAt
	///////////////////////////////////////

	updatedProject.ID = id
	return u.projectRepo.Update(id, updatedProject)
}

func (u *projectUseCase) Delete(user *models.User, id uint) error {
	// Check not needed but added to handle a not found error because gorm do not return
	// error if delete on a row that does not exist
	_, err := u.GetById(user, id)
	if err != nil {
		return err
	}

	return u.projectRepo.Delete(id)
}
