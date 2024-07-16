package models

type Document struct {
	GormModel
	Name     string `json:"name" gorm:"column:name"`
	Path     string `json:"path" gorm:"column:path"`
	SchoolId uint   `json:"schoolId" gorm:"column:school_id"`
	UserId   uint   `json:"userId" gorm:"column:user_id"`
	CourseId *uint  `json:"-" gorm:"column:course_id"`
	Course   Course `json:"course" gorm:"foreignKey:course_id;references:ID"`
}

type DocumentCreate struct {
	Name     string
	Byte     []byte
	CourseId *uint
}
