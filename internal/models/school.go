package models

type School struct {
	GormModel
	Name   string `json:"name" gorm:"column:name"`
	UserID uint   `gorm:"column:userId"`
}

type SchoolCreate struct {
	Name string `json:"name" binding:"required" validate:"min=1,max=64"`
}

type SchoolUpdate struct {
}
