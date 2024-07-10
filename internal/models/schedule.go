package models

type SignatureKind int

const (
	SIGNATURE_STUDENT       = 0
	SIGNATURE_TEACHER       = 1
	SIGNATURE_ADMINISTRATOR = 2
)

type Schedule struct {
	GormModel
	Time          uint   `json:"time" gorm:"column:time"`
	Duration      uint   `json:"duration" gorm:"column:time"`
	CourseId      uint   `json:"courseId" gorm:"column:course"`
	CampusId      uint   `json:"campusId" gorm:"column:campus"`
	ClassId       uint   `json:"classId" gorm:"column:class"`
	SignatureCode string `json:"-" gorm:"column:code"`
}

type ScheduleSignature struct {
	GormModel
	StudentId  uint          `json:"studentId" gorm:"column:student_id"`
	Student    User          `json:"student" gorm:"foreignKey:StudentId;references:ID"`
	ScheduleId uint          `json:"scheduleId" gorm:"column:schedule_id"`
	Schedule   Schedule      `json:"schedule" gorm:"foreignKey:ScheduleId;references:ID"`
	Kind       SignatureKind `json:"kind" gorm:"column:kind"`
}

type ScheduleSignatureCreate struct {
	SignatureCode string `json:"code"`
}

type ScheduleSignatureCode struct {
	SignatureCode string `json:"code"`
}

type ScheduleCreate struct {
	Time     *uint `json:"time" binding:"required"`
	Duration *uint `json:"duration" binding:"required"`
	CourseId *uint `json:"courseId" binding:"required" `
	CampusId *uint `json:"campusId" binding:"required"`
	ClassId  *uint `json:"classId" binding:"required"`
}

type ScheduleUpdate struct {
	Time     *uint `json:"time" binding:"required"`
	Duration *uint `json:"duration" binding:"required"`
	CourseId *uint `json:"courseId" binding:"required" `
	CampusId *uint `json:"campusId" binding:"required"`
	ClassId  *uint `json:"classId" binding:"required"`
}
