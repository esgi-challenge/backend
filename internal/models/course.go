package models

type Course struct {
	GormModel
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	TeacherId   uint   `json:"teacherId" gorm:"column:teacherId"`
	PathId      uint   `json:"pathId" gorm:"column:pathId"`
}

type CourseCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	TeacherId   *uint  `json:"teacherId" binding:"required"`
	PathId      *uint  `json:"pathId" binding:"required"`
}

type CourseUpdate struct {
	Description string `json:"description" binding:"required"`
	TeacherId   *uint  `json:"teacherId" binding:"required"`
	PathId      *uint  `json:"pathId" binding:"required"`
}
