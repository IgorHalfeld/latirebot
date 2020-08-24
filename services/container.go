package services

import "github.com/igorhalfeld/latirebot/structs"

type ProviderServiceInterface interface {
	GetProducts(clothingType structs.ClothingEnum) ([]structs.Product, error)
}

type Container struct {
	UserService      UserInterface
	RiachueloService ProviderServiceInterface
	RennerService    ProviderServiceInterface
	TelegramService  TelegramInterface
}
