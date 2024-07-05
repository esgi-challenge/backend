package models

type Campus struct {
	GormModel
	Name     string `json:"name" gorm:"column:name"`
	Location string `json:"location" gorm:"column:location"`
  Latitude float64 `json:"latitude" gorm:"column:latitude"`
  Longitude float64 `json:"longitude" gorm:"column:longitude"`
	SchoolId uint   `json:"schoolId" gorm:"column:school_id"`
}

type CampusCreate struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
}

type CampusUpdate struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
}
