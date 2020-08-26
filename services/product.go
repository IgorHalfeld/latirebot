package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/igorhalfeld/latirebot/repositories"
	"github.com/igorhalfeld/latirebot/structs"
)

type ProductService struct {
	repos repositories.Container
}

func NewProductService(repos repositories.Container) ProductService {
	return ProductService{repos}
}

func (ps ProductService) Create(ctx context.Context, product structs.Product) error {
	product.ID = uuid.New()
	err := ps.repos.ProductRepository.Create(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
