package models

type Course struct {
	GormModel
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	SchoolId    uint   `json:"schoolId" gorm:"column:school_id"`
	TeacherId   uint   `json:"teacherId" gorm:"column:teacher_id"`
	PathId      uint   `json:"pathId" gorm:"column:pathId"`
	Teacher     User   `json:"teacher" gorm:"foreignKey:TeacherId;references:ID"`
	Path        Path   `json:"path" gorm:"foreignKey:PathId;references:ID"`
}

type CourseCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	TeacherId   *uint  `json:"teacherId" binding:"required"`
	PathId      *uint  `json:"pathId" binding:"required"`
}

type CourseUpdate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	TeacherId   *uint  `json:"teacherId" binding:"required"`
	PathId      *uint  `json:"pathId" binding:"required"`
}
