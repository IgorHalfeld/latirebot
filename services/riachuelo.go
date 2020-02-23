package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/igorhalfeld/latirebot/structs"
)

// RiachueloService model
type RiachueloService struct{}

// NewRiachueloService creates a new instance
func NewRiachueloService() *RiachueloService {
	return &RiachueloService{}
}

// GetProducts get all products from riachuelo API
func (t *RiachueloService) GetProducts() ([]structs.Product, error) {
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

	products := []structs.Product{}
	json.Unmarshal(body, &products)

	return products, nil
}
