package service

import "Gin-Mongo/model"

type UserService interface {
	CreateUser(*model.User) error
	GetUser(*string) (*model.User, error)
	GetAll() ([]*model.User, error)
}
