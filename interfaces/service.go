package interfaces

import "github.com/igorhalfeld/latirebot/structs"

// IProductService represents a product service interface
type IProductService interface {
	GetProducts() ([]structs.Product, error)
}

// IAlertService  represents a product service interface
type IAlertService interface {
	Send(p structs.SendPayload)
}
