package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/igorhalfeld/latirebot/structs"
)

// RennerService model
type RennerService struct{}

// NewRennerService creates a new instance
func NewRennerService() *RennerService {
	return &RennerService{}
}

// GetProducts get all products from riachuelo API
func (rs *RennerService) GetProducts(clothingType structs.ClothingEnum) ([]structs.Product, error) {
	var products []structs.Product
	c := colly.NewCollector()

	URLs := rs.getURLWithAGivenClothingType(clothingType)

	for _, URL := range URLs {
		c.OnHTML(".wrapper.cf.results-list.js-results-list", func(e *colly.HTMLElement) {
			log.Println("Getting elements")

			e.ForEach("div.item_product", func(_ int, children *colly.HTMLElement) {
				var item structs.RennerCardItem

				json.Unmarshal([]byte(children.Attr("data-product-gtm")), &item)

				product, err := rs.getProductDetail(item.ID)
				if err != nil {
					log.Fatalln("Error on get product detail")
				}

				product.Image = "https:" + children.ChildAttr(".js-prod-link .js-images img", "src")
				product.Provider = "Renner"
				product.Link = "http:" + item.URL

				products = append(products, product)
			})
		})

		c.Visit(URL)
	}

	return products, nil
}

func (rs *RennerService) getProductDetail(id string) (structs.Product, error) {
	URL := "https://www.lojasrenner.com.br/rest/model/lrsa/api/CatalogActor/productBoxDataDesk?productId=" + id

	response, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
		return structs.Product{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
		return structs.Product{}, err
	}

	responseJSON := structs.RennerResponse{}
	json.Unmarshal(body, &responseJSON)

	product := structs.Product{
		Name:          responseJSON.Name,
		NormalPrice:   responseJSON.NormalPrice,
		DiscountPrice: responseJSON.DiscountPrice,
	}

	return product, nil
}

func (rs *RennerService) getURLWithAGivenClothingType(clothingType structs.ClothingEnum) []string {
	const (
		femaleURL string = "https://www.lojasrenner.com.br/c/feminino/-/N-4zo6za/p1"
		maleURL   string = "https://www.lojasrenner.com.br/c/masculino/-/N-1xeiyoy/p1"
	)
	var URL []string

	switch clothingType {
	case structs.ClothingTypeFemale:
		URL = []string{femaleURL}
	case structs.ClothingTypeMale:
		URL = []string{maleURL}
	case structs.ClothingTypeBoth:
		return []string{maleURL, femaleURL}
	default:
	}

	return URL
}

func formatPrice(raw string) float64 {
	if raw == "" {
		return 0
	}
	rawPrice := strings.Split(raw, " ")[1]
	value := strings.Replace(rawPrice, ",", ".", 2)
	price, _ := strconv.ParseFloat(value, 64)
	return price
}
