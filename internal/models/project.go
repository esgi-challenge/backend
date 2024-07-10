package models

type Project struct {
	GormModel
	Title         string `json:"title" gorm:"column:title"`
	EndDate       string `json:"endDate" gorm:"column:end_date"`
	CourseId      uint   `json:"courseId" gorm:"column:course_id"`
	ClassId       uint   `json:"classId" gorm:"column:class_id"`
	DocumentId    uint   `json:"documentId" gorm:"column:document_id"`
	GroupCapacity uint   `json:"groupCapacity" gorm:"column:group_capacity"`
}

type ProjectCreate struct {
	Title         string `json:"title" binding:"required" validate:"min=2,max=64"`
	EndDate       string `json:"endDate" binding:"required"`
	CourseId      *uint  `json:"courseId" binding:"required"`
	ClassId       *uint  `json:"classId" binding:"required"`
	DocumentId    *uint  `json:"documentId" binding:"required"`
	GroupCapacity *uint  `json:"groupCapacity" binding:"required"`
}

type ProjectUpdate struct {
	Title         string `json:"title" binding:"required" validate:"min=2,max=64"`
	EndDate       string `json:"endDate" binding:"required"`
	CourseId      *uint  `json:"courseId" binding:"required"`
	ClassId       *uint  `json:"classId" binding:"required"`
	DocumentId    *uint  `json:"documentId" binding:"required"`
	GroupCapacity *uint  `json:"groupCapacity" binding:"required"`
}

type ProjectStudent struct {
	GormModel
	Group     uint `json:"group" gorm:"column:group"`
	ProjectId uint `json:"projectId" gorm:"column:project_id"`
	StudentId uint `json:"studentId" gorm:"column:student_id"`
}

type ProjectStudentCreate struct {
	Group *uint `json:"group" binding:"required"`
}
