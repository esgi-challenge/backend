package models

type Informations struct {
	GormModel
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	SchoolId    uint   `json:"schoolId" gorm:"column:school_id"`
}

type InformationsCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
