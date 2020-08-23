package handlers

import "net/http"

type HealthInterface interface {
	Check(w http.ResponseWriter, _ *http.Request)
}

type ProductsHandler interface {
	Look()
}

type Container struct {
	HealthHandler   HealthInterface
	ProductsHandler ProductsHandler
}
