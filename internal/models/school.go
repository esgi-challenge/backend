package models

type School struct {
	GormModel
	Name   string `json:"name" gorm:"column:name"`
	UserID uint   `gorm:"column:userId"`
}

type SchoolCreate struct {
	Name string `json:"name" binding:"required" validate:"min=1,max=64"`
}

type SchoolInvite struct {
	SchoolId  uint
	Firstname string `json:"firstname" binding:"required" validate:"min=1,max=128"`
	Lastname  string `json:"lastname" binding:"required" validate:"min=1,max=128"`
	Email     string `json:"email" binding:"required" validate:"min=1,max=128"`
}

type SchoolUpdate struct {
}
