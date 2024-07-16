package usecase

import (
	"fmt"
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

	project, err = u.projectRepo.Create(project)
	if err != nil {
		return nil, err
	}

	//Retrieve preload
	return u.projectRepo.GetPreloadById(project.ID)
}

func (u *projectUseCase) JoinProject(user *models.User, join *models.ProjectStudentCreate, id uint) (*models.ProjectStudent, error) {
	project, err := u.GetById(user, id)

	fmt.Println(project)

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

	fmt.Println(*join.Group)

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
	if *user.UserKind == models.TEACHER {
		return u.projectRepo.GetAllByTeacher(user)
	} else {
		return u.projectRepo.GetAll(user)
	}
}

func (u *projectUseCase) GetById(user *models.User, id uint) (*models.Project, error) {
	return u.projectRepo.GetById(user, id)
}

func (u *projectUseCase) GetGroups(user *models.User, id uint) (*[]models.ProjectGroup, error) {
	maxGroup := uint(1)
	groups := []models.ProjectGroup{}

	students, err := u.projectRepo.GetGroups(user, id)

	if err != nil {
		return nil, err
	}

	for _, student := range *students {
		if student.Group >= maxGroup {
			maxGroup = student.Group + 1
		}

		var existingGroup *models.ProjectGroup

		for _, group := range groups {
			if group.GroupId == student.Group {
				existingGroup = &group
			}
		}

		if existingGroup == nil {
			groups = append(groups, models.ProjectGroup{
				GroupId: student.Group,
				Users: []models.User{
					student.Student,
				},
			})
		} else {
			existingGroup.Users = append(existingGroup.Users, student.Student)
		}

	}

	groups = append(groups, models.ProjectGroup{
		GroupId: maxGroup,
		Users:   []models.User{},
	})

	return &groups, nil
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
	project, err := u.projectRepo.Update(id, updatedProject)
	if err != nil {
		return nil, err
	}

	return u.projectRepo.GetPreloadById(project.ID)
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
