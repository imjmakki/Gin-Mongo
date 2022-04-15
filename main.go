package main

import (
	"Gin-Mongo/controller"
	"Gin-Mongo/service"
	"github.com/gin-gonic/gin"
)

var (
	server         *gin.Engine
	userservice    service.UserService
	usercontroller controller.UserController
)

func main() {

}
