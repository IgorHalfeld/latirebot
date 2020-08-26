package services

import (
	"context"

	"github.com/igorhalfeld/latirebot/structs"
)

type ProviderServiceInterface interface {
	GetProducts(clothingType structs.ClothingEnum) ([]structs.Product, error)
}

type ProductInterface interface {
	Create(ctx context.Context, product structs.Product) error
}

type UserInterface interface {
	ReadAll(ctx context.Context) ([]structs.User, error)
}

type TelegramInterface interface {
	ListenMessages()
	SendNotification(payload structs.NotificationPayload)
}

type Container struct {
	UserService      UserInterface
	ProductService   ProductInterface
	RiachueloService ProviderServiceInterface
	RennerService    ProviderServiceInterface
	TelegramService  TelegramInterface
}
