package main

import (
	"Gin-Mongo/controller"
	"Gin-Mongo/service"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	server         *gin.Engine
	userservice    service.UserService
	usercontroller controller.UserController
	ctx            context.Context
	usercollectoin *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	mongocon := options.Client().ApplyURI("mongodb://localhost:27017")
}

func main() {

}
