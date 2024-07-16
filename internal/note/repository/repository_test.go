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
	repo := NewNoteRepository(db)

	note := &models.Note{
		Title:       "title",
		Description: "description",
	}

	t.Run("Create", func(t *testing.T) {
		mock.ExpectBegin()
		// Should be ExpectExec because of INSERT query, but ExpectQuery needed (known issue)
		// https://github.com/DATA-DOG/go-sqlmock/issues/118
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, note.Title, note.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(note.ID, note.Title, note.Description))
		mock.ExpectCommit()

		createdNote, err := repo.Create(note)

		assert.NoError(t, err)
		assert.NotNil(t, createdNote)
		assert.Equal(t, createdNote, note)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Create Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, note.Title, note.Description).
			WillReturnError(errors.New("create failed"))
		mock.ExpectRollback()

		createdNote, err := repo.Create(note)

		assert.Error(t, err)
		assert.Nil(t, createdNote)
		assert.EqualError(t, err, "create failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewNoteRepository(db)

	note1 := models.Note{
		Title:       "title1",
		Description: "description1",
	}

	note2 := models.Note{
		Title:       "title2",
		Description: "description2",
	}

	t.Run("Get All", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(note1.ID, note1.Title, note1.Description).
			AddRow(note2.ID, note2.Title, note2.Description)

		mock.ExpectQuery(getAllQuery).
			WillReturnRows(rows)

		notes, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, notes)
		assert.Len(t, *notes, 2)
		assert.Equal(t, &note1, &(*notes)[0])
		assert.Equal(t, &note2, &(*notes)[1])
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get All Error", func(t *testing.T) {
		mock.ExpectQuery(getAllQuery).
			WillReturnError(errors.New("get all failed"))

		notes, err := repo.GetAll()

		assert.Error(t, err)
		assert.Nil(t, notes)
		assert.EqualError(t, err, "get all failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewNoteRepository(db)

	note := models.Note{
		Title:       "title",
		Description: "description",
	}

	t.Run("Get By Id", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(note.ID, note.Title, note.Description)

		mock.ExpectQuery(getQuery).
			WillReturnRows(row)

		dbNote, err := repo.GetById(note.ID)

		assert.NoError(t, err)
		assert.NotNil(t, dbNote)
		assert.Equal(t, dbNote, &note)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get By Id Error", func(t *testing.T) {
		mock.ExpectQuery(getQuery).
			WillReturnError(errors.New("get failed"))

		note, err := repo.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, note)
		assert.EqualError(t, err, "get failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewNoteRepository(db)

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
	repo := NewNoteRepository(db)

	note := &models.Note{
		Title:       "title",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, note.Title, note.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(note.ID, note.Title, note.Description))
		mock.ExpectCommit()

		updatedNote, err := repo.Update(1, note)

		assert.NoError(t, err)
		assert.NotNil(t, updatedNote)
		assert.Equal(t, updatedNote, note)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, note.Title, note.Description).
			WillReturnError(errors.New("update failed"))
		mock.ExpectRollback()

		updatedNote, err := repo.Create(note)

		assert.Error(t, err)
		assert.Nil(t, updatedNote)
		assert.EqualError(t, err, "update failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
