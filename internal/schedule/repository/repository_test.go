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
	repo := NewScheduleRepository(db)

	schedule := &models.Schedule{
		Title:       "title",
		Description: "description",
	}

	t.Run("Create", func(t *testing.T) {
		mock.ExpectBegin()
		// Should be ExpectExec because of INSERT query, but ExpectQuery needed (known issue)
		// https://github.com/DATA-DOG/go-sqlmock/issues/118
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, schedule.Title, schedule.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(schedule.ID, schedule.Title, schedule.Description))
		mock.ExpectCommit()

		createdSchedule, err := repo.Create(schedule)

		assert.NoError(t, err)
		assert.NotNil(t, createdSchedule)
		assert.Equal(t, createdSchedule, schedule)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Create Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, schedule.Title, schedule.Description).
			WillReturnError(errors.New("create failed"))
		mock.ExpectRollback()

		createdSchedule, err := repo.Create(schedule)

		assert.Error(t, err)
		assert.Nil(t, createdSchedule)
		assert.EqualError(t, err, "create failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewScheduleRepository(db)

	schedule1 := models.Schedule{
		Title:       "title1",
		Description: "description1",
	}

	schedule2 := models.Schedule{
		Title:       "title2",
		Description: "description2",
	}

	t.Run("Get All", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(schedule1.ID, schedule1.Title, schedule1.Description).
			AddRow(schedule2.ID, schedule2.Title, schedule2.Description)

		mock.ExpectQuery(getAllQuery).
			WillReturnRows(rows)

		schedules, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, schedules)
		assert.Len(t, *schedules, 2)
		assert.Equal(t, &schedule1, &(*schedules)[0])
		assert.Equal(t, &schedule2, &(*schedules)[1])
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get All Error", func(t *testing.T) {
		mock.ExpectQuery(getAllQuery).
			WillReturnError(errors.New("get all failed"))

		schedules, err := repo.GetAll()

		assert.Error(t, err)
		assert.Nil(t, schedules)
		assert.EqualError(t, err, "get all failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewScheduleRepository(db)

	schedule := models.Schedule{
		Title:       "title",
		Description: "description",
	}

	t.Run("Get By Id", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(schedule.ID, schedule.Title, schedule.Description)

		mock.ExpectQuery(getQuery).
			WillReturnRows(row)

		dbSchedule, err := repo.GetById(schedule.ID)

		assert.NoError(t, err)
		assert.NotNil(t, dbSchedule)
		assert.Equal(t, dbSchedule, &schedule)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get By Id Error", func(t *testing.T) {
		mock.ExpectQuery(getQuery).
			WillReturnError(errors.New("get failed"))

		schedule, err := repo.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, schedule)
		assert.EqualError(t, err, "get failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewScheduleRepository(db)

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
	repo := NewScheduleRepository(db)

	schedule := &models.Schedule{
		Title:       "title",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, schedule.Title, schedule.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(schedule.ID, schedule.Title, schedule.Description))
		mock.ExpectCommit()

		updatedSchedule, err := repo.Update(1, schedule)

		assert.NoError(t, err)
		assert.NotNil(t, updatedSchedule)
		assert.Equal(t, updatedSchedule, schedule)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, schedule.Title, schedule.Description).
			WillReturnError(errors.New("update failed"))
		mock.ExpectRollback()

		updatedSchedule, err := repo.Create(schedule)

		assert.Error(t, err)
		assert.Nil(t, updatedSchedule)
		assert.EqualError(t, err, "update failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
