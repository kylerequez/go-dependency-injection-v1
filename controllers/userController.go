package controllers

import (
	"github.com/kylerequez/go-dependency-injection/services"
)

type Controller interface {
	GetAllUsers()
}

type UserController struct {
	us *services.UserService
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		us: us,
	}
}

func (uc *UserController) GetAllUsers() {
	uc.us.GetAllUsers()
}
