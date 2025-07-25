package dto

type CreateUserDto struct {
	Code         string
	Firstname    string
	Lastname     string
	Nickname     string
	Age          int
	Birthdate    string
	Phone        string
	Role         string
	ProfileImage []byte
	Username     string
	Password     string
}