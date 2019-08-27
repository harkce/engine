package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Router decide which request goes to which handler
func Router() http.Handler {
	router := httprouter.New()

	router.HandlerFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})
	return router
}
