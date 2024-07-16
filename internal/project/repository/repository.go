package repository

import (
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/project"
	"gorm.io/gorm"
)

type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) project.Repository {
	return &projectRepo{db: db}
}

func (r *projectRepo) Create(project *models.Project) (*models.Project, error) {
	if err := r.db.Create(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (r *projectRepo) GetPreloadById(id uint) (*models.Project, error) {
	var project models.Project

	if err := r.db.Model(&models.Project{}).Preload("Class").Preload("Course").Preload("Document").Where("id = ?", id).Find(&project).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *projectRepo) GetAllByTeacher(user *models.User) (*[]models.Project, error) {
	var projects []models.Project

	if err := r.db.Model(&models.Project{}).Preload("Class").Preload("Course").Preload("Document").Where("teacher_id = ?", user.ID).Find(&projects).Error; err != nil {
		return nil, err
	}

	return &projects, nil
}

func (r *projectRepo) GetAll(user *models.User) (*[]models.Project, error) {
	var projects []models.Project

	if err := r.db.Model(&models.Project{}).Joins("Course").Preload("Class").Joins("left join classes on classes.id = projects.class_id").Preload("Course.Teacher").Joins("left join schools on classes.school_id = schools.id").Where("classes.id = ?", user.ClassRefer).Or("schools.user_id = ?", user.ID).Find(&projects).Error; err != nil {
		return nil, err
	}

	return &projects, nil
}

func (r *projectRepo) GetJoined(projectId uint, userId uint) (*[]models.ProjectStudent, error) {
	var projectStudent []models.ProjectStudent

	if err := r.db.Raw(checkIfJoined, userId, projectId).Scan(&projectStudent).Error; err != nil {
		return nil, err
	}

	return &projectStudent, nil
}

func (r *projectRepo) DeleteJoined(projectId uint, userId uint) error {
	var projectStudent models.ProjectStudent

	if err := r.db.Raw(checkIfJoinedUniq, userId, projectId).Scan(&projectStudent).Error; err != nil {
		return err
	}

	if err := r.db.Debug().Delete(&models.ProjectStudent{}, projectStudent.ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *projectRepo) JoinProject(project *models.ProjectStudent) (*models.ProjectStudent, error) {
	if err := r.db.Create(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (r *projectRepo) GetGroups(user *models.User, id uint) (*[]models.ProjectStudent, error) {
	var project *[]models.ProjectStudent

	if err := r.db.Model(&models.ProjectStudent{}).Joins("Student").Joins("Student.Class").Preload("Project").Where("project_id = ?", id).Find(&project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (r *projectRepo) GetById(user *models.User, id uint) (*models.Project, error) {
	var project models.Project

	if err := r.db.Model(&models.Project{}).Joins("Course").Preload("Class").Joins("left join classes on classes.id = projects.class_id").Preload("Course.Teacher").Joins("left join schools on classes.school_id = schools.id").Where("projects.id = ?", id).Where("classes.id = ?", user.ClassRefer).Or("schools.user_id = ?", user.ID).Find(&project).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *projectRepo) Update(id uint, project *models.Project) (*models.Project, error) {
	if err := r.db.Save(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (r *projectRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.Project{}, id).Error; err != nil {
		return err
	}

	return nil
}
