package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/igorhalfeld/latirebot/structs"
)

// RiachueloService model
type RiachueloService struct{}

// NewRiachueloService creates a new instance
func NewRiachueloService() RiachueloService {
	return RiachueloService{}
}

// GetProducts get all products from riachuelo API
func (rs RiachueloService) GetProducts(clothingType structs.ClothingEnum) ([]structs.Product, error) {
	var products []structs.Product
	URLs := rs.getURLWithAGivenClothingType(clothingType)

	for _, URL := range URLs {
		response, err := http.Get(URL)
		log.Println("[riachuelo] Getting elements")
		if err != nil {
			log.Fatalln("[riachuelo]", err)
			return []structs.Product{}, err
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln("[riachuelo]", err)
			return []structs.Product{}, err
		}

		responseJSON := []structs.RiachueloResponse{}
		json.Unmarshal(body, &responseJSON)

		for _, product := range responseJSON {
			responseData, _ := json.Marshal(product)
			np, _ := strconv.ParseFloat(product.MinPrice01, 64)
			dp, _ := strconv.ParseFloat(product.ChMaxPrice01, 64)

			products = append(products, structs.Product{
				SKU:           product.SKU,
				Name:          product.Name,
				NormalPrice:   np,
				DiscountPrice: dp,
				Provider:      "Riachuelo",
				Link:          "https://www.riachuelo.com.br/" + product.URLKey,
				ImageURL:      product.SmallImage,
				ResponseData:  string(responseData),
			})
		}
	}

	return products, nil
}

func (rs RiachueloService) getURLWithAGivenClothingType(clothingType structs.ClothingEnum) []string {
	const (
		maleURL   string = "https://www.riachuelo.com.br/elasticsearch/data/products?category_id=63"
		femaleURL string = "https://www.riachuelo.com.br/elasticsearch/data/products?category_id=18"
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
