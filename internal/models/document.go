package models

type Document struct {
	GormModel
	Path   string `json:"path" gorm:"column:path"`
	UserId uint   `json:"description" gorm:"column:description"`
}

type DocumentCreate struct {
	Byte []byte
}

type DocumentUpdate struct {
}
