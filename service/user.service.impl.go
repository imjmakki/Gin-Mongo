package service

import "go.mongodb.org/mongo-driver/mongo"

type UserServiceImpl struct {
	usercollection *mongo.Collection
}
