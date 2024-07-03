package repository

import (
	"fmt"

	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/internal/user"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) (*models.User, error) {
	userDb, _ := r.GetByEmail(user.Email)

	if userDb != nil {
		return nil, gorm.ErrDuplicatedKey
	}

	if err := r.db.Create(user).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}

	return user, nil
}

func (r *userRepo) GetAll() (*[]models.User, error) {
	var users []models.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepo) GetById(id uint) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if result := r.db.First(&user, "email = ?", email); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepo) GetByResetCode(resetCode string) (*models.User, error) {
	user := &models.User{}
	if result := r.db.First(&user, "password_reset_code = ?", resetCode); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepo) GetByInvitationCode(invitationCode string) (*models.User, error) {
	user := &models.User{}
	if result := r.db.First(&user, "invitation_code = ?", invitationCode); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepo) Update(id uint, user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) Delete(id uint) error {
	if err := r.db.Debug().Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
