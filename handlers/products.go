package handlers

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"

	"github.com/igorhalfeld/latirebot/services"
	"github.com/igorhalfeld/latirebot/structs"
)

// Scan looking for products
type Product struct {
	services services.Container
}

// NewScanHandler creates a new instance
func NewProductHandler(services services.Container) *Product {
	return &Product{services}
}

func (p Product) Look() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Println("Riachelo dispatched")
		products, err := p.services.RiachueloService.GetProducts()
		if err != nil {
			log.Println(err)
		}
		p.comparePricesAndSendAlert(products)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		log.Println("Renner dispatched")
		products, err := p.services.RennerService.GetProducts()
		if err != nil {
			log.Println(err)
		}
		p.comparePricesAndSendAlert(products)
		wg.Done()
	}()

	wg.Wait()
}

func (p Product) comparePricesAndSendAlert(products []structs.Product) {
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
			p.services.TelegramService.Send(structs.SendPayload{
				Name:      product.Name,
				Provider:  product.Provider,
				Caption:   p.buildCaption(product, dP, nP),
				ChatID:    "158277392",
				Photo:     product.Image,
				ParseMode: "HTML",
			})
		}
	}

	log.Println("All sent!")
}

func (p *Product) buildCaption(product structs.Product, cV float64, nV float64) string {
	var value float64 = cV

	price := value / 1

	return "<strong>" + strings.ToUpper(product.Name) + "</strong> ⚡️ " + strings.ToUpper(product.Provider) + " ⚡️ " +
		"<i>R$" + fmt.Sprintf("%.2f", price) + "</i> - <s>R$" + fmt.Sprintf("%.2f", nV) + "</s> ⚡️ " +
		"<a href='" + product.Link + "'>Ver detalhes 💵</a>"
}

func sortDiscount(products []structs.Product) []structs.Product {
	sort.SliceStable(products, func(i, j int) bool {
		return (products[i].NormalPrice / products[i].DiscountPrice) > (products[j].NormalPrice / products[j].DiscountPrice) // Less funciton reversed for one-shot reversing
	})

	return products
}
