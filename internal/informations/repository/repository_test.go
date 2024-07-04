package repository

import (
	"errors"
	"testing"

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

func TestCreate(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewInformationsRepository(db)

	informations := &models.Informations{
		Title:       "title",
		Description: "description",
	}

	t.Run("Create", func(t *testing.T) {
		mock.ExpectBegin()
		// Should be ExpectExec because of INSERT query, but ExpectQuery needed (known issue)
		// https://github.com/DATA-DOG/go-sqlmock/issues/118
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, informations.Title, informations.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(informations.ID, informations.Title, informations.Description))
		mock.ExpectCommit()

		createdInformations, err := repo.Create(informations)

		assert.NoError(t, err)
		assert.NotNil(t, createdInformations)
		assert.Equal(t, createdInformations, informations)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Create Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, informations.Title, informations.Description).
			WillReturnError(errors.New("create failed"))
		mock.ExpectRollback()

		createdInformations, err := repo.Create(informations)

		assert.Error(t, err)
		assert.Nil(t, createdInformations)
		assert.EqualError(t, err, "create failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewInformationsRepository(db)

	informations1 := models.Informations{
		Title:       "title1",
		Description: "description1",
	}

	informations2 := models.Informations{
		Title:       "title2",
		Description: "description2",
	}

	t.Run("Get All", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(informations1.ID, informations1.Title, informations1.Description).
			AddRow(informations2.ID, informations2.Title, informations2.Description)

		mock.ExpectQuery(getAllQuery).
			WillReturnRows(rows)

		informationss, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, informationss)
		assert.Len(t, *informationss, 2)
		assert.Equal(t, &informations1, &(*informationss)[0])
		assert.Equal(t, &informations2, &(*informationss)[1])
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get All Error", func(t *testing.T) {
		mock.ExpectQuery(getAllQuery).
			WillReturnError(errors.New("get all failed"))

		informationss, err := repo.GetAll()

		assert.Error(t, err)
		assert.Nil(t, informationss)
		assert.EqualError(t, err, "get all failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewInformationsRepository(db)

	informations := models.Informations{
		Title:       "title",
		Description: "description",
	}

	t.Run("Get By Id", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(informations.ID, informations.Title, informations.Description)

		mock.ExpectQuery(getQuery).
			WillReturnRows(row)

		dbInformations, err := repo.GetById(informations.ID)

		assert.NoError(t, err)
		assert.NotNil(t, dbInformations)
		assert.Equal(t, dbInformations, &informations)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get By Id Error", func(t *testing.T) {
		mock.ExpectQuery(getQuery).
			WillReturnError(errors.New("get failed"))

		informations, err := repo.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, informations)
		assert.EqualError(t, err, "get failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewInformationsRepository(db)

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

func TestUpdate(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewInformationsRepository(db)

	informations := &models.Informations{
		Title:       "title",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, informations.Title, informations.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(informations.ID, informations.Title, informations.Description))
		mock.ExpectCommit()

		updatedInformations, err := repo.Update(1, informations)

		assert.NoError(t, err)
		assert.NotNil(t, updatedInformations)
		assert.Equal(t, updatedInformations, informations)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, informations.Title, informations.Description).
			WillReturnError(errors.New("update failed"))
		mock.ExpectRollback()

		updatedInformations, err := repo.Create(informations)

		assert.Error(t, err)
		assert.Nil(t, updatedInformations)
		assert.EqualError(t, err, "update failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
