package services

import "github.com/igorhalfeld/latirebot/structs"

type ProviderServiceInterface interface {
	GetProducts() ([]structs.Product, error)
}

type AlertServiceInterface interface {
	Send(p structs.SendPayload)
}

type Container struct {
	UserService      UserInterface
	RiachueloService ProviderServiceInterface
	RennerService    ProviderServiceInterface
	TelegramService  AlertServiceInterface
}
