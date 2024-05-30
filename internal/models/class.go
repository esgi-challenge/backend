package models

type Class struct {
	GormModel
	Name     string `json:"name" gorm:"column:name"`
	PathId   uint   `json:"pathId" gorm:"column:pathId"`
	Students []User `json:"students" gorm:"foreignKey:ClassRefer"`
}

type ClassCreate struct {
	Name   string `json:"name" binding:"required"`
	PathId uint   `json:"pathId" binding:"required"`
}

type ClassAdd struct {
	// Do like this to allow zero value
	UserId *uint `json:"userId" binding:"required" `
}

type ClassUpdate struct {
	Name   string `json:"name" binding:"required"`
	PathId uint   `json:"pathId" binding:"required"`
}
