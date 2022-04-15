package controller

import (
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

func (u *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (u *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}
