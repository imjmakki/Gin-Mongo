package service

import (
	"Gin-Mongo/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *model.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*model.User, error) {
	var user *model.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*model.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *model.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}
