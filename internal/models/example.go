package models

type Example struct {
	GormModel
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

type ExampleCreate struct {
	Title       string `json:"title" binding:"required" validate:"min=8,max=64"`
	Description string `json:"description" binding:"required"`
}

type ExampleUpdate struct {
	Title       string `json:"title" binding:"required" validate:"min=8,max=64"`
	Description string `json:"description" binding:"required"`
}
