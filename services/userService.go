package services

import (
	"fmt"

	"github.com/kylerequez/go-dependency-injection/repository"
)

type Services interface {
	GetAllUsers()
}

type UserService struct {
	ur *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (us *UserService) GetAllUsers() {
	// Business Code here

	users, err := us.ur.GetAllUsers()
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println("RESULT", users)

}
