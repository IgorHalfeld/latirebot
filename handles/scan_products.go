package handles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/igorhalfeld/latirebot/services"
	"github.com/igorhalfeld/latirebot/structs"
)

// Scan looking for products
type Scan struct {
	RiachueloService *services.RiachueloService
	TelegramService  *services.TelegramService
}

// NewScanHandler creates a new instance
func NewScanHandler(s Scan) *Scan {
	return &s
}

// Look starts looking for products
func (s *Scan) Look() {
	products, err := s.RiachueloService.GetProducts()
	if err != nil {
		fmt.Println(err)
	}

	for _, product := range products[0:100] {
		cV, _ := strconv.ParseFloat(product.ChMaxPrice01, 64)
		nV, _ := strconv.ParseFloat(product.MinPrice01, 64)

		if cV != 0 && cV < nV {
			s.TelegramService.Send(structs.SendPayload{
				Name:      product.Name,
				Caption:   buildCaption(product, cV, nV),
				ChatID:    "158277392",
				Photo:     product.SmallImage,
				ParseMode: "HTML",
			})
		}
	}

	fmt.Println("All sent!")
}

func buildCaption(product structs.Product, cV float64, nV float64) string {
	var value float64 = cV

	price := value / 1

	return "<strong>" + strings.ToUpper(product.Name) + "</strong> ⚡️ " +
		"<i>R$" + fmt.Sprintf("%.2f", price) + "</i> - <s>R$" + fmt.Sprintf("%.2f", nV) + "</s> ⚡️ " +
		"<a href='https://www.riachuelo.com.br/" + product.URLKey + "'>Ver detalhes💵</a>"
}
