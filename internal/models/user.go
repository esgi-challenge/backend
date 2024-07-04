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
	Firstname         string   `json:"firstname" gorm:"column:firstname"`
	Lastname          string   `json:"lastname" gorm:"column:lastname"`
	Email             string   `json:"email" gorm:"column:email;unique;not null"`
	Password          string   `json:"password" gorm:"column:password"`
	InvitationCode    string   `json:"invitationCode" gorm:"column:invitation_code"`
	PasswordResetCode string   `json:"passwordResetCode" gorm:"column:password_reset_code"`
	UserKind          UserKind `json:"userKind" gorm:"column:user_kind"`
	SchoolId          *uint    `json:"schoolId" gorm:"column:school_id"`
	ClassRefer        *uint    `json:"classRefer" gorm:"column:class_refer"`
}

type UserCreate struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	UserKind  UserKind `json:"userKind"`
	SchoolId  *uint    `json:"schoolId"`
}

type SchoolUserCreate struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type SchoolUserUpdate struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type SendResetMail struct {
	Email string `json:"email" binding:"required"`
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
