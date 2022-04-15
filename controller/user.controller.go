package controller

import (
	"Gin-Mongo/model"
	"Gin-Mongo/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func New(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (u *UserServiceImpl) GetUser(name *string) {
	ctx.JSON(200, "")
}

func (u *UserServiceImpl) GetAll() {
	ctx.JSON(200, "")
}

func (u *UserServiceImpl) UpdateUser(user *model.User) {
	ctx.JSON(200, "")
}

func (u *UserServiceImpl) DeleteUser(name *string) {
	ctx.JSON(200, "")
}
