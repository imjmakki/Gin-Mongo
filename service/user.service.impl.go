package service

import (
	"Gin-Mongo/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func (u *UserServiceImpl) CreateUser(user *model.User) error {
	return nil
}

func (u *UserServiceImpl) GetUser(name *string) (*model.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) GetAll() []*model.User {
	return nil
}

func (u *UserServiceImpl) UpdateUser(user *model.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}
