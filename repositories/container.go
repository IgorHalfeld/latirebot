package repositories

import (
	"context"

	"github.com/igorhalfeld/latirebot/structs"
)

type UserInterface interface {
	Create(context context.Context, user structs.User) error
	ReadOneByUsername(context context.Context) (structs.User, error)
	ReadAll(context context.Context) ([]structs.User, error)
}

type ProductInterface interface {
	Create(context context.Context, product structs.Product) error
}

type Container struct {
	UserRepository    UserInterface
	ProductRepository ProductInterface
}
