package repositories

import (
	"context"

	"github.com/igorhalfeld/latirebot/structs"
)

type UserInterface interface {
	Create(context context.Context, ID string) error
	ReadAll(context context.Context) ([]structs.Users, error)
}

type Container struct {
	UserRepository UserInterface
}
