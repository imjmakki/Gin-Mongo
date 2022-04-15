package controller

import (
	"Gin-Mongo/model"
	"Gin-Mongo/service"
)

type UserController struct {
	UserService service.UserService
}

func New(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (u *UserServiceImpl) CreateUser(user *model.User) {
	return nil
}

func (u *UserServiceImpl) GetUser(name *string) {
	return nil, nil
}

func (u *UserServiceImpl) GetAll() {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *model.User) {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) {
	return nil
}
