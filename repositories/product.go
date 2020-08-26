package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/igorhalfeld/latirebot/structs"
	"github.com/jmoiron/sqlx"
)

var (
	errorProductCreate  = errors.New("Product create failed")
	errorProductReadAll = errors.New("Product readall failed")
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (pr ProductRepository) Create(ctx context.Context, product structs.Product) error {
	query := `
		INSERT INTO
			products(
				id,
				sku,
				name,
				provider,
				response_data,
				normal_price,
				discount_price
			)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
	`
	result, err := pr.db.ExecContext(ctx, query, product.ID, product.GetSKU(), product.Name, product.Provider, product.ResponseData, product.NormalPrice, product.DiscountPrice)
	if err != nil {
		log.Println(err)
		return errorProductCreate
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Println(err)
		return errorProductCreate
	}

	log.Println("product creating", product.ID)
	return nil
}
