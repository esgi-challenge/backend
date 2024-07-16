package models

type Note struct {
	GormModel
	Value     int     `json:"value" gorm:"column:value"`
	StudentId uint    `json:"studentId" gorm:"column:student_id"`
	TeacherId uint    `json:"teacherId" gorm:"column:teacher_id"`
	ProjectId uint    `json:"projectId" gorm:"column:project_id"`
	Student   User    `json:"student" gorm:"foreignKey:StudentId;references:ID"`
	Teacher   User    `json:"teacher" gorm:"foreignKey:TeacherId;references:ID"`
	Project   Project `json:"project" gorm:"foreignKey:ProjectId;references:ID"`
}

type NoteCreate struct {
	Value     int  `json:"value" binding:"required"`
	ProjectId uint `json:"projectId" binding:"required"`
	StudentId uint `json:"studentId" binding:"required"`
}

type NoteUpdate struct {
	Value     int  `json:"value" binding:"required"`
	ProjectId uint `json:"projectId" binding:"required"`
	StudentId uint `json:"studentId" binding:"required"`
}
