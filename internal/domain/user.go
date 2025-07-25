package domain

import (
	"coop_student_backend/internal/dto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Code string
	Firstname string
	Lastname string
	Nickname string
	Age int 
	Birthdate string
	Phone string
	Role string
	AliveStatus bool
	EducationStatus bool
	GovermmentStatus bool
	ProfileImage []byte
	Order []Order `gorm:"foreignKey:UserID"`
	UserLoginID uint
	UserLogin UserLogin `gorm:"foreignKey:UserLoginID;constraint:OnDelete:CASCADE;"`
}



type UserService interface{
	FindAll(uid int)(*[]User,error)
	FindById(id string) (*User, error)
	Create(createUserDto dto.CreateUserDto) (*User,error)
	Update(id string,updateProductDto dto.UpdateUserDto) (*User,error)
	Delete(id string) (*User,error)
}