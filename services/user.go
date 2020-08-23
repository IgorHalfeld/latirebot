package services

import "github.com/igorhalfeld/latirebot/repositories"

type UserInterface interface {
}

type UserService struct {
	repos repositories.Container
}

func NewUserService(repos repositories.Container) *UserService {
	return &UserService{}
}
