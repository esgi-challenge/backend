package models

type Schedule struct {
	GormModel
	Time     string `json:"time" gorm:"column:time"`
	CourseId uint   `json:"courseId" gorm:"column:course"`
	CampusId uint   `json:"campusId" gorm:"column:campus"`
	ClassId  uint   `json:"classId" gorm:"column:class"`
}

type ScheduleCreate struct {
	Time     string `json:"time" binding:"required"`
	CourseId *uint  `json:"courseId" binding:"required" `
	CampusId *uint  `json:"campusId" binding:"required"`
	ClassId  *uint  `json:"classId" binding:"required"`
}

type ScheduleUpdate struct {
	Time     string `json:"time" binding:"required"`
	CourseId *uint  `json:"courseId" binding:"required" `
	CampusId *uint  `json:"campusId" binding:"required"`
	ClassId  *uint  `json:"classId" binding:"required"`
}
