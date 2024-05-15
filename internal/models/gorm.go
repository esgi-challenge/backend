package models

import (
	"time"

	"gorm.io/gorm"
)

// Custom gorm.Model to add lowercase to json
type GormModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt" gorm:"<-:create"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
