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

func (uc *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUser(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/get-users", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete", uc.DeleteUser)
}
