package models

import "golang.org/x/crypto/bcrypt"

type UserKind int

const (
	STUDENT       = 0
	TEACHER       = 1
	ADMINISTRATOR = 2
	SUPERADMIN    = 3
)

type User struct {
	GormModel
	Firstname      string   `json:"firstname" gorm:"column:firstname"`
	Lastname       string   `json:"lastname" gorm:"column:lastname"`
	Email          string   `json:"email" gorm:"column:email"`
	Password       string   `json:"password" gorm:"column:password"`
	InvitationCode string   `json:"invitationCode" gorm:"column:invitationCode"`
	UserKind       UserKind `json:"userKind" gorm:"column:userKind"`
	SchoolId       *uint    `json:"schoolId" gorm:"column:schoolId"`
	ClassRefer     *uint    `json:"classRefer"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
