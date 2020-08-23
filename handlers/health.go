package handlers

import (
	"fmt"
	"net/http"
)

type Health struct{}

func NewHealthHandler() Health {
	return Health{}
}

func (s Health) Check(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Ok")
}
