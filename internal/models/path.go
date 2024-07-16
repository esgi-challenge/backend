package models

type Path struct {
	GormModel
	ShortName string `json:"shortName" gorm:"column:short_name"`
	School    School `json:"school" gorm:"foreignKey:SchoolId;references:ID""`
	SchoolId  uint   `json:"schoolId" gorm:"column:school_id"`
	LongName  string `json:"longName" gorm:"column:long_name"`
}

type PathCreate struct {
	ShortName string `json:"shortName" binding:"required"`
	LongName  string `json:"longName"`
}

type PathUpdate struct {
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
}
