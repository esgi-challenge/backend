package models

type Class struct {
	GormModel
	Name     string `json:"name" gorm:"column:name"`
	SchoolId uint   `json:"schoolId" gorm:"column:school_id"`
	PathId   uint   `json:"pathId" gorm:"column:path_id"`
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

type ClassRemove struct {
	UserId *uint `json:"userId" binding:"required" `
}

type ClassUpdate struct {
	Name   string `json:"name" binding:"required"`
	PathId uint   `json:"pathId" binding:"required"`
}
