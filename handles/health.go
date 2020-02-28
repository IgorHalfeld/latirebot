package handles

import (
	"fmt"
	"net/http"
)

// Health says that is OK
type Health struct{}

// NewHealthHandler creates a new instance
func NewHealthHandler() *Health {
	return &Health{}
}

// Check say something when ping on 80 port
func (s *Health) Check() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Ok")
	})

	http.ListenAndServe(":80", nil)
}
