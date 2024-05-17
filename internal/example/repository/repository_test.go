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
	repo := NewExampleRepository(db)

  example := &models.Example{
    Title:       "title",
    Description: "description",
  }

  t.Run("Create", func(t *testing.T) {
    mock.ExpectBegin()
    // Should be ExpectExec because of INSERT query, but ExpectQuery needed (known issue)
    // https://github.com/DATA-DOG/go-sqlmock/issues/118
    mock.ExpectQuery(createQuery).
      WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, example.Title, example.Description).
      WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(example.ID, example.Title, example.Description))
    mock.ExpectCommit()

    createdExample, err := repo.Create(example)

    assert.NoError(t, err)
    assert.NotNil(t, createdExample)
    assert.Equal(t, createdExample, example)
    assert.NoError(t, mock.ExpectationsWereMet())
  })

  t.Run("Create Error", func(t *testing.T) {
    mock.ExpectBegin()
    mock.ExpectQuery(createQuery).
      WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, example.Title, example.Description).
      WillReturnError(errors.New("create failed"))
    mock.ExpectRollback()

    createdExample, err := repo.Create(example)

    assert.Error(t, err)
    assert.Nil(t, createdExample)
    assert.EqualError(t, err, "create failed")
    assert.NoError(t, mock.ExpectationsWereMet())
  })
}

func TestGetAll(t *testing.T) {
  t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewExampleRepository(db)

  example1 := models.Example{
    Title: "title1",  
    Description: "description1",
  }

  example2 := models.Example{
    Title: "title2",  
    Description: "description2",
  }

  t.Run("Get All", func(t *testing.T) {
    rows := sqlmock.NewRows([]string{"id", "title", "description"}).
      AddRow(example1.ID, example1.Title, example1.Description).
      AddRow(example2.ID, example2.Title, example2.Description)

    mock.ExpectQuery(getAllQuery).
      WillReturnRows(rows)

    examples, err := repo.GetAll()

    assert.NoError(t, err)
    assert.NotNil(t, examples)
    assert.Len(t, *examples, 2)
    assert.Equal(t, &example1, &(*examples)[0])
    assert.Equal(t, &example2, &(*examples)[1])
    assert.NoError(t, mock.ExpectationsWereMet())
  })

  t.Run("Get All Error", func(t *testing.T) {
    mock.ExpectQuery(getAllQuery).
      WillReturnError(errors.New("get all failed"))

    examples, err := repo.GetAll()

    assert.Error(t, err)
    assert.Nil(t, examples)
    assert.EqualError(t, err, "get all failed")
    assert.NoError(t, mock.ExpectationsWereMet())
  })
}

func TestGetById(t *testing.T) {
  t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewExampleRepository(db)

  example := models.Example{
    Title: "title",  
    Description: "description",
  }

  t.Run("Get By Id", func(t *testing.T) {
    row := sqlmock.NewRows([]string{"id", "title", "description"}).
      AddRow(example.ID, example.Title, example.Description)

    mock.ExpectQuery(getQuery).
      WillReturnRows(row)

    dbExample, err := repo.GetById(example.ID)

    assert.NoError(t, err)
    assert.NotNil(t, dbExample)
    assert.Equal(t, dbExample, &example)
    assert.NoError(t, mock.ExpectationsWereMet())
  })

  t.Run("Get By Id Error", func(t *testing.T) {
    mock.ExpectQuery(getQuery).
      WillReturnError(errors.New("get failed"))

    example, err := repo.GetById(1)

    assert.Error(t, err)
    assert.Nil(t, example)
    assert.EqualError(t, err, "get failed")
    assert.NoError(t, mock.ExpectationsWereMet())
  })
}

func TestDelete(t *testing.T) {
  t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewExampleRepository(db)

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
	repo := NewExampleRepository(db)

  example := &models.Example{
    Title: "title",  
    Description: "description",
  }

  t.Run("Update", func(t *testing.T) {
    mock.ExpectBegin()
    mock.ExpectQuery(updateQuery).
      WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, example.Title, example.Description).
      WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(example.ID, example.Title, example.Description))
    mock.ExpectCommit()

    updatedExample, err := repo.Update(1, example)

    assert.NoError(t, err)
    assert.NotNil(t, updatedExample)
    assert.Equal(t, updatedExample, example)
    assert.NoError(t, mock.ExpectationsWereMet())
  })

  t.Run("Update Error", func(t *testing.T) {
    mock.ExpectBegin()
    mock.ExpectQuery(updateQuery).
      WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, example.Title, example.Description).
      WillReturnError(errors.New("update failed"))
    mock.ExpectRollback()

    updatedExample, err := repo.Create(example)

    assert.Error(t, err)
    assert.Nil(t, updatedExample)
    assert.EqualError(t, err, "update failed")
    assert.NoError(t, mock.ExpectationsWereMet())
  })
}
