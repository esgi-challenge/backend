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
	repo := NewCampusRepository(db)

	campus1 := models.Campus{
		Name: "name",
		Location:  "location",
		Latitude:  1,
		Longitude:  1,
	}

	campus2 := models.Campus{
		Name: "name2",
		Location:  "location2",
		Latitude:  1,
		Longitude:  1,
	}

	t.Run("Get All", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "location", "latitude", "longitude"}).
			AddRow(campus1.ID, campus1.Name, campus1.Location, campus1.Latitude, campus1.Longitude).
			AddRow(campus2.ID, campus2.Name, campus2.Location, campus2.Latitude, campus2.Longitude)

		mock.ExpectQuery(getAllQuery).
			WillReturnRows(rows)

		campuses, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, campuses)
		assert.Len(t, *campuses, 2)
		assert.Equal(t, &campus1, &(*campuses)[0])
		assert.Equal(t, &campus2, &(*campuses)[1])
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get All Error", func(t *testing.T) {
		mock.ExpectQuery(getAllQuery).
			WillReturnError(errors.New("get all failed"))

		campuss, err := repo.GetAll()

		assert.Error(t, err)
		assert.Nil(t, campuss)
		assert.EqualError(t, err, "get all failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewCampusRepository(db)

  campus := models.Campus{
    Name: "name",
    Location:  "location",
    Latitude:  1,
    Longitude:  1,
  }

	t.Run("Get By Id", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "name", "location", "latitude", "longitude"}).
			AddRow(campus.ID, campus.Name, campus.Location, campus.Latitude, campus.Longitude)

		mock.ExpectQuery(getQuery).
			WillReturnRows(row)

		dbcampus, err := repo.GetById(campus.ID)

		assert.NoError(t, err)
		assert.NotNil(t, dbcampus)
		assert.Equal(t, dbcampus, &campus)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get By Id Error", func(t *testing.T) {
		mock.ExpectQuery(getQuery).
			WillReturnError(errors.New("get failed"))

		campus, err := repo.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, campus)
		assert.EqualError(t, err, "get failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewCampusRepository(db)

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
