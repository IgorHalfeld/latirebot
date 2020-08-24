package handlers

import (
	"context"
	"fmt"
	"log"
	"sort"
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
	ctx := context.Background()
	users, _ := p.services.UserService.ReadAll(ctx)

	log.Println("Total of users to be notified:", len(users))

	for _, user := range users {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			log.Println("Riachelo dispatched")
			fmt.Println("ClothingType", user)
			products, err := p.services.RiachueloService.GetProducts(user.ClothingType)
			if err != nil {
				log.Println(err)
			}
			p.comparePricesAndSendAlert(products, user)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			log.Println("Renner dispatched")
			products, err := p.services.RennerService.GetProducts(user.ClothingType)
			if err != nil {
				log.Println(err)
			}
			p.comparePricesAndSendAlert(products, user)
			wg.Done()
		}()

		wg.Wait()
	}
}

func (p Product) comparePricesAndSendAlert(products []structs.Product, user structs.User) {
	var lastIndex int = 10
	// Guarantee to send the better discounts
	products = sortDiscount(products)

	for _, product := range products[0:lastIndex] {

		if product.DiscountPrice != 0 && product.DiscountPrice < product.NormalPrice {
			payload := structs.NotificationPayload{
				Product:       product,
				User:          user,
				DiscountPrice: product.DiscountPrice,
				NormalPrice:   product.NormalPrice,
			}

			p.services.TelegramService.SendNotification(payload)
		}
	}

	log.Println("All sent!")
}

func sortDiscount(products []structs.Product) []structs.Product {
	sort.SliceStable(products, func(i, j int) bool {
		// Less function reversed for one-shot reversing
		return (products[i].NormalPrice / products[i].DiscountPrice) >
			(products[j].NormalPrice / products[j].DiscountPrice)
	})

	return products
}
