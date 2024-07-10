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
	repo := NewChatRepository(db)

	chat := &models.Chat{
		Title:       "title",
		Description: "description",
	}

	t.Run("Create", func(t *testing.T) {
		mock.ExpectBegin()
		// Should be ExpectExec because of INSERT query, but ExpectQuery needed (known issue)
		// https://github.com/DATA-DOG/go-sqlmock/issues/118
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, chat.Title, chat.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(chat.ID, chat.Title, chat.Description))
		mock.ExpectCommit()

		createdChat, err := repo.Create(chat)

		assert.NoError(t, err)
		assert.NotNil(t, createdChat)
		assert.Equal(t, createdChat, chat)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Create Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(createQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, chat.Title, chat.Description).
			WillReturnError(errors.New("create failed"))
		mock.ExpectRollback()

		createdChat, err := repo.Create(chat)

		assert.Error(t, err)
		assert.Nil(t, createdChat)
		assert.EqualError(t, err, "create failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewChatRepository(db)

	chat1 := models.Chat{
		Title:       "title1",
		Description: "description1",
	}

	chat2 := models.Chat{
		Title:       "title2",
		Description: "description2",
	}

	t.Run("Get All", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(chat1.ID, chat1.Title, chat1.Description).
			AddRow(chat2.ID, chat2.Title, chat2.Description)

		mock.ExpectQuery(getAllQuery).
			WillReturnRows(rows)

		chats, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, chats)
		assert.Len(t, *chats, 2)
		assert.Equal(t, &chat1, &(*chats)[0])
		assert.Equal(t, &chat2, &(*chats)[1])
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get All Error", func(t *testing.T) {
		mock.ExpectQuery(getAllQuery).
			WillReturnError(errors.New("get all failed"))

		chats, err := repo.GetAll()

		assert.Error(t, err)
		assert.Nil(t, chats)
		assert.EqualError(t, err, "get all failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewChatRepository(db)

	chat := models.Chat{
		Title:       "title",
		Description: "description",
	}

	t.Run("Get By Id", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(chat.ID, chat.Title, chat.Description)

		mock.ExpectQuery(getQuery).
			WillReturnRows(row)

		dbChat, err := repo.GetById(chat.ID)

		assert.NoError(t, err)
		assert.NotNil(t, dbChat)
		assert.Equal(t, dbChat, &chat)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get By Id Error", func(t *testing.T) {
		mock.ExpectQuery(getQuery).
			WillReturnError(errors.New("get failed"))

		chat, err := repo.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, chat)
		assert.EqualError(t, err, "get failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	db, mock := setupMockDB(t)
	repo := NewChatRepository(db)

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
	repo := NewChatRepository(db)

	chat := &models.Chat{
		Title:       "title",
		Description: "description",
	}

	t.Run("Update", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, chat.Title, chat.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "comment"}).AddRow(chat.ID, chat.Title, chat.Description))
		mock.ExpectCommit()

		updatedChat, err := repo.Update(1, chat)

		assert.NoError(t, err)
		assert.NotNil(t, updatedChat)
		assert.Equal(t, updatedChat, chat)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(updateQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, chat.Title, chat.Description).
			WillReturnError(errors.New("update failed"))
		mock.ExpectRollback()

		updatedChat, err := repo.Create(chat)

		assert.Error(t, err)
		assert.Nil(t, updatedChat)
		assert.EqualError(t, err, "update failed")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
