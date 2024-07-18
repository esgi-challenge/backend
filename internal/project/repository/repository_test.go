package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	return gormDB, mock
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewProjectRepository(db)

	project1 := models.Project{
		Title:   "project",
		EndDate: time.Now(),
	}

	project2 := models.Project{
		Title:   "project2",
		EndDate: time.Now(),
	}

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}

	t.Run("Get All", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "end_date"}).
			AddRow(project1.ID, project1.Title, project1.EndDate).
			AddRow(project2.ID, project2.Title, project2.EndDate)

		mock.ExpectQuery(getAllQuery).
			WillReturnRows(rows)

		projects, err := repo.GetAll(adminUser)

		assert.NoError(t, err)
		assert.NotNil(t, projects)
		assert.Len(t, *projects, 2)
		assert.Equal(t, &project1, &(*projects)[0])
		assert.Equal(t, &project2, &(*projects)[1])
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get All Error", func(t *testing.T) {
		mock.ExpectQuery(getAllQuery).
			WillReturnError(errors.New("get all failed"))

		projects, err := repo.GetAll(adminUser)

		assert.Error(t, err)
		assert.Nil(t, projects)
		assert.EqualError(t, err, "get all failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewProjectRepository(db)

	project := models.Project{
		Title:   "title",
		EndDate: time.Now(),
	}

	adminUser := &models.User{
		UserKind: models.NewUserKind(models.ADMINISTRATOR),
	}

	t.Run("Get By Id", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "title", "end_date"}).
			AddRow(project.ID, project.Title, project.EndDate)

		mock.ExpectQuery(getQuery).
			WillReturnRows(row)

		dbProject, err := repo.GetById(adminUser, project.ID)

		assert.NoError(t, err)
		assert.NotNil(t, dbProject)
		assert.Equal(t, dbProject, &project)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get By Id Error", func(t *testing.T) {
		mock.ExpectQuery(getQuery).
			WillReturnError(errors.New("get failed"))

		project, err := repo.GetById(adminUser, 1)

		assert.Error(t, err)
		assert.Nil(t, project)
		assert.EqualError(t, err, "get failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewProjectRepository(db)

	t.Run("Delete", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(deleteQuery).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Delete(1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("delete Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(deleteQuery).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnError(errors.New("delete failed"))
		mock.ExpectRollback()

		err := repo.Delete(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "delete failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
