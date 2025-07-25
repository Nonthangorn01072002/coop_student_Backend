package services

import (
	"coop_student_backend/internal/domain"
	"coop_student_backend/internal/dto"

	"gorm.io/gorm"
)

type UserService struct {
	db               *gorm.DB
	userLoginService *UserLoginService 
}

func NewUserService(db *gorm.DB, loginService *UserLoginService) *UserService {
	return &UserService{
		db:               db,
		userLoginService: loginService, 
	}
}

func (u *UserService) FindAll(uid int) (*[]domain.User, error) {
	var chkAccessUser domain.User

	if err := u.db.First(&chkAccessUser, "id = ?", uid).Error; err != nil {
		return nil, err
	}

	if chkAccessUser.Role != "TEACHER" {
		return nil, nil
	}

	var users []domain.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}


func (u *UserService) FindById(id string) (*domain.User, error){
	var user domain.User
	if err := u.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) Create(createUserDto dto.CreateUserDto) (*domain.User, error) {

	createLogin, err := u.userLoginService.Create(createUserDto.Username,createUserDto.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Code:             createUserDto.Code,
		Firstname:        createUserDto.Firstname,
		Lastname:         createUserDto.Lastname,
		Nickname:         createUserDto.Nickname,
		Age:              createUserDto.Age,
		Birthdate:        createUserDto.Birthdate,
		Phone:            createUserDto.Phone,
		Role:             createUserDto.Role,
		AliveStatus:      true,
		EducationStatus:  true,
		GovermmentStatus: false,
		ProfileImage:     createUserDto.ProfileImage,
		UserLoginID:      createLogin.ID,
	}

	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) Update(id string,updateUserDto dto.UpdateUserDto) (*domain.User,error){
	var currentUser domain.User
	if err := u.db.First(&currentUser, "id=?", id).Error; err != nil{
		return nil, err
	}
	currentUser.Code = updateUserDto.Code
	currentUser.Firstname = updateUserDto.Firstname
	currentUser.Lastname = updateUserDto.Lastname
	currentUser.Nickname = updateUserDto.Nickname
	currentUser.Age = updateUserDto.Age
	currentUser.Birthdate = updateUserDto.Birthdate
	currentUser.Phone = updateUserDto.Phone
	currentUser.Role = updateUserDto.Role
	if err := u.db.Save(currentUser).Error; err != nil{
		return nil, err
	}
	return &currentUser, nil
}

func (u *UserService) Delete(id string) (*domain.User,error){
	var user domain.User
	if err := u.db.Delete(&user, "id=?", id).Error; err != nil{
		return nil, err
	}
	return &user,nil
}
