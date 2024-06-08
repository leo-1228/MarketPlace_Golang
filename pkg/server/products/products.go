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

// ProductContext handles loading product item with given ID.
// Returns default Empty response if product not found.
func ProductContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var product *Product
		if chi.URLParam(req, "productId") == "" {
			render.Respond(w, req, httputils.NotFoundResponse)
			return
		}

		product = &Product{
			ID:       1,
			Name:     "Toaster",
			Price:    19.99,
			Quantity: 100,
		}

		ctx := context.WithValue(req.Context(), productContextKey, product)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
