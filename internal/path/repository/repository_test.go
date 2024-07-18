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

func TestGetAll(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewPathRepository(db)

	path1 := models.Path{
		ShortName: "name",
		LongName:  "name",
	}

	path2 := models.Path{
		ShortName: "name",
		LongName:  "name",
	}

	t.Run("Get All", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "short_name", "long_name"}).
			AddRow(path1.ID, path1.ShortName, path1.LongName).
			AddRow(path2.ID, path2.ShortName, path2.LongName)

		mock.ExpectQuery(getAllQuery).
			WillReturnRows(rows)

		paths, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, paths)
		assert.Len(t, *paths, 2)
		assert.Equal(t, &path1, &(*paths)[0])
		assert.Equal(t, &path2, &(*paths)[1])
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get All Error", func(t *testing.T) {
		mock.ExpectQuery(getAllQuery).
			WillReturnError(errors.New("get all failed"))

		paths, err := repo.GetAll()

		assert.Error(t, err)
		assert.Nil(t, paths)
		assert.EqualError(t, err, "get all failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewPathRepository(db)

	path := models.Path{
		ShortName: "name",
		LongName:  "name",
	}

	t.Run("Get By Id", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "short_name", "long_name"}).
			AddRow(path.ID, path.ShortName, path.LongName)

		mock.ExpectQuery(getQuery).
			WillReturnRows(row)

		dbPath, err := repo.GetById(path.ID)

		assert.NoError(t, err)
		assert.NotNil(t, dbPath)
		assert.Equal(t, dbPath, &path)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get By Id Error", func(t *testing.T) {
		mock.ExpectQuery(getQuery).
			WillReturnError(errors.New("get failed"))

		path, err := repo.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, path)
		assert.EqualError(t, err, "get failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewPathRepository(db)

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
