package models

import (
	"gorm.io/gorm"
)

type UserKind int

const (
	STUDENT       = 0
	TEACHER       = 1
	ADMINISTRATOR = 2
	SUPERADMIN    = 3
)

type User struct {
	gorm.Model
	Firstname string   `gorm:"column:firstname"`
	Lastname  string   `gorm:"column:lastname"`
	Email     string   `gorm:"column:email"`
	Password  string   `gorm:"column:password"`
	UserKind  UserKind `gorm:"column:userKind"`
}
