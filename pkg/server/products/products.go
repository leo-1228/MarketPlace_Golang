package products

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Product dto object for a product .
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type contextKey string

var productContextKey = contextKey("product")

// RegisterRoutes applies the routes for the products service.
func RegisterRoutes(router chi.Router) http.Handler {
	router.Get("/", ListProducts)
	router.Post("/", CreateProducts)

	router.Route("/{productId}", func(r chi.Router) {
		r.Use(ProductContext)
		r.Get("/", GetProductItem)
		r.Patch("/", UpdateProductItem)
		r.Delete("/", DeleteProductItem)
	})

	return router
}
