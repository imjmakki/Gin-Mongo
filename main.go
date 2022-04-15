package main

import (
	"Gin-Mongo/controller"
	"Gin-Mongo/service"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server         *gin.Engine
	userservice    service.UserService
	usercontroller controller.UserController
	ctx            context.Context
	usercollectoin *mongo.Collection
	mongoclient    *mongo.Client
)

func main() {

}
