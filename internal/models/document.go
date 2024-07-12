package models

type Document struct {
	GormModel
	Path     string `json:"path" gorm:"column:path"`
	UserId   uint   `json:"description" gorm:"column:description"`
	CourseId *uint  `json:"-" gorm:"column:course_id"`
	Course   Course `json:"course" gorm:"foreignKey:course_id;references:ID"`
}

type DocumentCreate struct {
	Byte     []byte
	CourseId *uint
}

type DocumentUpdate struct {
}
