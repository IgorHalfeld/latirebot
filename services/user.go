package services

import (
	"context"

	"github.com/igorhalfeld/latirebot/repositories"
	"github.com/igorhalfeld/latirebot/structs"
)

type UserService struct {
	repos repositories.Container
}

func NewUserService(repos repositories.Container) *UserService {
	return &UserService{repos}
}

func (us UserService) ReadAll(ctx context.Context) ([]structs.User, error) {
	users, err := us.repos.UserRepository.ReadAll(ctx)
	if err != nil {
		return []structs.User{}, err
	}

	return users, nil
}
