package main

import (
	"github.com/go-chi/chi"
	"github.com/lowkruc/rataid-product-golang.git/internal/handler"
)

// newRoutes is function to initial Chi Route handler, with handler parameter
func newRoutes(handler handler.HTTPHandler) *chi.Mux {
	// initial chi.Mux
	router := chi.NewRouter()

	// add group endpoint to `/v1`
	router.Route("/v1", func(v1 chi.Router) {

		// set handler `/v1/product`
		v1.Get("/products", handler.GetProduct)
	})

	return router
}
