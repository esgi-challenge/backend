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
	Duration      uint   `json:"duration" gorm:"column:duration"`
	CourseId      uint   `json:"courseId" gorm:"column:course"`
	CampusId      uint   `json:"campusId" gorm:"column:campus"`
	ClassId       uint   `json:"classId" gorm:"column:class"`
	SchoolId      uint   `json:"schoolId" gorm:"column:school_id"`
	QrCodeEnabled bool   `json:"qrCodeEnabled" gorm:"column:qr_code_enabled"`
	SignatureCode string `json:"-" gorm:"column:code"`
	Course        Course `json:"course" gorm:"foreignKey:CourseId;references:ID"`
	Campus        Campus `json:"campus" gorm:"foreignKey:CampusId;references:ID"`
	Class         Class  `json:"class" gorm:"foreignKey:ClassId;references:ID"`
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
	UserId        uint   `json:"userId"`
}

type ScheduleSignatureCode struct {
	SignatureCode string `json:"code"`
}

type ScheduleCreate struct {
	Time          *uint `json:"time" binding:"required"`
	QrCodeEnabled bool  `json:"qrCodeEnabled" binding:"required"`
	Duration      *uint `json:"duration" binding:"required"`
	CourseId      *uint `json:"courseId" binding:"required"`
	CampusId      *uint `json:"campusId" binding:"required"`
	ClassId       *uint `json:"classId" binding:"required"`
}

type ScheduleUpdate struct {
	Time          *uint `json:"time" binding:"required"`
	QrCodeEnabled bool  `json:"qrCodeEnabled" binding:"required"`
	Duration      *uint `json:"duration" binding:"required"`
	CourseId      *uint `json:"courseId" binding:"required" `
	CampusId      *uint `json:"campusId" binding:"required"`
	ClassId       *uint `json:"classId" binding:"required"`
}

type ScheduleGet struct {
	Schedule Schedule `json:"schedule" binding:"required"`
	Course   Course   `json:"course" binding:"required"`
	Campus   Campus   `json:"campus" binding:"required"`
}

type ScheduleSignatureGet struct {
	Students  []User              `json:"students" binding:"required"`
	Signature []ScheduleSignature `json:"signatures" binding:"required"`
}
