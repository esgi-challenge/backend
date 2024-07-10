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
	repo := NewDocumentRepository(db)

	document := &models.Document{
		Title:       "title",
		Description: "description",
	}

	t.Run("Create", func(t *testing.T) {
		mock.ExpectBegin()
		// Should be ExpectExec because of INSERT query, but ExpectQuery needed (known issue)
		// https://github.com/DATA-DOG/go-sqlmock/issues/118
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, document.Title, document.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(document.ID, document.Title, document.Description))
		mock.ExpectCommit()

		createdDocument, err := repo.Create(document)

		assert.NoError(t, err)
		assert.NotNil(t, createdDocument)
		assert.Equal(t, createdDocument, document)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Create Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, document.Title, document.Description).
			WillReturnError(errors.New("create failed"))
		mock.ExpectRollback()

		createdDocument, err := repo.Create(document)

		assert.Error(t, err)
		assert.Nil(t, createdDocument)
		assert.EqualError(t, err, "create failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewDocumentRepository(db)

	document1 := models.Document{
		Title:       "title1",
		Description: "description1",
	}

	document2 := models.Document{
		Title:       "title2",
		Description: "description2",
	}

	t.Run("Get All", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(document1.ID, document1.Title, document1.Description).
			AddRow(document2.ID, document2.Title, document2.Description)

		mock.ExpectQuery(getAllQuery).
			WillReturnRows(rows)

		documents, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, documents)
		assert.Len(t, *documents, 2)
		assert.Equal(t, &document1, &(*documents)[0])
		assert.Equal(t, &document2, &(*documents)[1])
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get All Error", func(t *testing.T) {
		mock.ExpectQuery(getAllQuery).
			WillReturnError(errors.New("get all failed"))

		documents, err := repo.GetAll()

		assert.Error(t, err)
		assert.Nil(t, documents)
		assert.EqualError(t, err, "get all failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewDocumentRepository(db)

	document := models.Document{
		Title:       "title",
		Description: "description",
	}

	t.Run("Get By Id", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(document.ID, document.Title, document.Description)

		mock.ExpectQuery(getQuery).
			WillReturnRows(row)

		dbDocument, err := repo.GetById(document.ID)

		assert.NoError(t, err)
		assert.NotNil(t, dbDocument)
		assert.Equal(t, dbDocument, &document)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get By Id Error", func(t *testing.T) {
		mock.ExpectQuery(getQuery).
			WillReturnError(errors.New("get failed"))

		document, err := repo.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, document)
		assert.EqualError(t, err, "get failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewDocumentRepository(db)

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
	repo := NewDocumentRepository(db)

	document := &models.Document{
		Title:       "title",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, document.Title, document.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(document.ID, document.Title, document.Description))
		mock.ExpectCommit()

		updatedDocument, err := repo.Update(1, document)

		assert.NoError(t, err)
		assert.NotNil(t, updatedDocument)
		assert.Equal(t, updatedDocument, document)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, document.Title, document.Description).
			WillReturnError(errors.New("update failed"))
		mock.ExpectRollback()

		updatedDocument, err := repo.Create(document)

		assert.Error(t, err)
		assert.Nil(t, updatedDocument)
		assert.EqualError(t, err, "update failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}