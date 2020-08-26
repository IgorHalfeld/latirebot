package structs

import (
	"time"
)

// RiachueloResponse is riachuelo json response
type RiachueloResponse struct {
	ChMaxPrice01   string    `json:"ch_max_price_0_1"`
	MinPrice01     string    `json:"min_price_0_1"`
	SmallImage     string    `json:"small_image"`
	ActivationDate time.Time `json:"activation_date"`
	SKU            string    `json:"sku"`
	MaxPrice01     string    `json:"max_price_0_1"`
	Marca          string    `json:"marca"`
	CorSimples     []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"cor_simples"`
	MarcaValue   string `json:"marca_value"`
	URLKey       string `json:"url_key"`
	ChMinPrice01 string `json:"ch_min_price_0_1"`
	Name         string `json:"name"`
}

// RennerResponse is renner json response
type RennerResponse struct {
	DiscountPrice float64 `json:"salePrice"`
	NormalPrice   float64 `json:"listPrice"`
	Name          string  `json:"displayName"`
}

// RennerCardItem is the payload on card HTML
type RennerCardItem struct {
	URL   string `json:"url"`
	ID    string `json:"id"`
	SkuID string `json:"sku"`
}
