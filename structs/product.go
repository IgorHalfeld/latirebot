package structs

import "github.com/google/uuid"

type ProductProviderEnum string

const (
	ProductProviderRiachualo ProductProviderEnum = "RIACHUELO"
	ProductProviderRenner    ProductProviderEnum = "RENNER"
)

type Product struct {
	ID            uuid.UUID           `db:"id"`
	SKU           string              `db:"sku"`
	Link          string              `db:"link"`
	Name          string              `db:"name"`
	Provider      ProductProviderEnum `db:"provider"`
	ImageURL      string              `db:"image_url"`
	NormalPrice   float64             `db:"normal_price"`
	DiscountPrice float64             `db:"discount_price"`
	ResponseData  string              `db:"response_data"`
}

func (p *Product) GetSKU() string {
	return string(p.Provider) + "-" + p.SKU
}
