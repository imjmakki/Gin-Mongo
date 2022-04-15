package main

import (
	"Gin-Mongo/controller"
	"Gin-Mongo/service"
	"context"
	"github.com/gin-gonic/gin"
)

var (
	server         *gin.Engine
	userservice    service.UserService
	usercontroller controller.UserController
	ctx            context.Context
)

func main() {

}
