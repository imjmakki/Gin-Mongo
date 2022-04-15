package controller

import "Gin-Mongo/service"

type UserController struct {
	UserService service.UserService
}

func New() UserController {
}
