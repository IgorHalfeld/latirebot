package handlers

import (
	"context"
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

	chUsers := make(chan structs.User, 1)
	go p.Dispatch(chUsers)

	for _, user := range users {
		chUsers <- user
	}

	close(chUsers)
}

func (p Product) Dispatch(chUsers chan structs.User) {
	select {
	case user, ok := <-chUsers:
		if ok {
			var wg sync.WaitGroup

			wg.Add(1)
			go func() {
				log.Println("Riachelo dispatched")
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
	default:
	}
}

func (p Product) comparePricesAndSendAlert(products []structs.Product, user structs.User) {
	var lastIndex int = 10

	for _, product := range products[0:lastIndex] {

		// create product without blocking
		go func() {
			_ = p.services.ProductService.Create(context.Background(), product)
		}()

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
