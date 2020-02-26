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
func NewRiachueloService() *RiachueloService {
	return &RiachueloService{}
}

// GetProducts get all products from riachuelo API
func (rs *RiachueloService) GetProducts() ([]structs.Product, error) {
	URL := "https://www.riachuelo.com.br/elasticsearch/data/products?category_id=63"

	response, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
		return []structs.Product{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
		return []structs.Product{}, err
	}

	responseJSON := []structs.RiachueloResponse{}
	json.Unmarshal(body, &responseJSON)

	var products []structs.Product

	for _, product := range responseJSON {
		nP, _ := strconv.ParseFloat(product.MinPrice01, 64)
		dP, _ := strconv.ParseFloat(product.ChMaxPrice01, 64)

		products = append(products, structs.Product{
			Provider:      "Riachuelo",
			Name:          product.Name,
			NormalPrice:   nP,
			DiscountPrice: dP,
			Link:          "https://www.riachuelo.com.br/" + product.URLKey,
			Image:         product.SmallImage,
		})
	}

	return products, nil
}
