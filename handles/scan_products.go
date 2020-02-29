package handles

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"

	"github.com/igorhalfeld/latirebot/interfaces"
	"github.com/igorhalfeld/latirebot/structs"
)

// Scan looking for products
type Scan struct {
	RiachueloService interfaces.IProductService
	TelegramService  interfaces.IAlertService
	RennerService    interfaces.IProductService
}

// NewScanHandler creates a new instance
func NewScanHandler(s Scan) *Scan {
	return &s
}

// Look starts looking for products
func (s *Scan) Look() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Println("Riachelo dispatched")
		products, err := s.RiachueloService.GetProducts()
		if err != nil {
			log.Println(err)
		}
		s.comparePricesAndSendAlert(products)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		log.Println("Renner dispatched")
		products, err := s.RennerService.GetProducts()
		if err != nil {
			log.Println(err)
		}
		s.comparePricesAndSendAlert(products)
		wg.Done()
	}()

	wg.Wait()
}

func (s *Scan) comparePricesAndSendAlert(products []structs.Product) {
	var lastIndex int = 100
	if len(products) < 100 {
		lastIndex = len(products) - 1
	} else {
		products = sortDiscount(products) // In case recieved more than 100 products, guarantee to send the better discounts
	}

	for _, product := range products[0:lastIndex] {
		dP := product.DiscountPrice
		nP := product.NormalPrice

		if dP != 0 && dP < nP {
			s.TelegramService.Send(structs.SendPayload{
				Name:      product.Name,
				Provider:  product.Provider,
				Caption:   s.buildCaption(product, dP, nP),
				ChatID:    "158277392",
				Photo:     product.Image,
				ParseMode: "HTML",
			})
		}
	}

	log.Println("All sent!")
}

func sortDiscount(products []structs.Product) []structs.Product {
	sort.SliceStable(products, func(i, j int) bool {
		return (products[i].NormalPrice / products[i].DiscountPrice) > (products[j].NormalPrice / products[j].DiscountPrice) // Less funciton reversed for one-shot reversing
	})

	return products
}

func (s *Scan) buildCaption(product structs.Product, cV float64, nV float64) string {
	var value float64 = cV

	price := value / 1

	return "<strong>" + strings.ToUpper(product.Name) + "</strong> ⚡️ " + strings.ToUpper(product.Provider) + " ⚡️ " +
		"<i>R$" + fmt.Sprintf("%.2f", price) + "</i> - <s>R$" + fmt.Sprintf("%.2f", nV) + "</s> ⚡️ " +
		"<a href='" + product.Link + "'>Ver detalhes 💵</a>"
}
