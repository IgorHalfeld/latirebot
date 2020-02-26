package structs

// Product represents a product standard
type Product struct {
	Name          string
	Image         string
	NormalPrice   float64
	DiscountPrice float64
	Provider      string
	Link          string
}
