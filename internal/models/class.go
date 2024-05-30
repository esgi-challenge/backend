package models

type Class struct {
	GormModel
	Name   string `json:"name" gorm:"column:name"`
	PathId uint   `json:"pathId" gorm:"column:pathId"`
}

type ClassCreate struct {
	Name   string `json:"name" binding:"required" validate:"min=8,max=64"`
	PathId uint   `json:"pathId" binding:"required"`
}

type ClassUpdate struct {
	Name   string `json:"name" binding:"required"`
	PathId uint   `json:"pathId" binding:"required"`
}
