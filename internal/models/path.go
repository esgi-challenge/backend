package models

type Path struct {
	GormModel
	Name        string `json:"name" gorm:"column:name"`
	SchoolId    uint   `json:"schoolId" gorm:"column:schoolId"`
	Description string `json:"description" gorm:"column:description"`
}

type PathCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	SchoolId    uint   `json:"schoolId" binding:"required"`
}

type PathUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SchoolId    uint   `json:"schoolId"`
}
