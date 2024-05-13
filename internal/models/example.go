package models

import (
	"gorm.io/gorm"
)

type Example struct {
  gorm.Model `json:"-"`
  Title       string `json:"title" gorm:"column:title"`
  Description string `json:"description" gorm:"column:description"`
}
