package models

type Campus struct {
	GormModel
	Name     string `json:"name" gorm:"column:name"`
	Location string `json:"location" gorm:"column:location"`
	Code     string `json:"code" gorm:"column:code"`
	SchoolId uint   `json:"schoolId" gorm:"column:schoolId"`
}

type CampusCreate struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
	Code     string `json:"code"`
	SchoolId uint   `json:"schoolId" binding:"required"`
}

type CampusUpdate struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
	Code     string `json:"code"`
	SchoolId uint   `json:"schoolId" binding:"required"`
}
