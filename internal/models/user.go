package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}
