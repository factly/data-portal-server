package action

import (
	"net/http"

	"github.com/factly/data-portal-api/action/cart"
	"github.com/factly/data-portal-api/action/cartitem"
	"github.com/factly/data-portal-api/action/category"
	"github.com/factly/data-portal-api/action/currency"
	"github.com/factly/data-portal-api/action/membership"
	"github.com/factly/data-portal-api/action/order"
	"github.com/factly/data-portal-api/action/payment"
	"github.com/factly/data-portal-api/action/plan"
	"github.com/factly/data-portal-api/action/product"
	"github.com/factly/data-portal-api/action/tag"
	_ "github.com/factly/data-portal-api/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterRoutes - CRUD servies
func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	//r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Mount("/currencies", currency.Router())

	r.Route("/users", func(r chi.Router) {
		r.Post("/", CreateUser)
		r.Get("/", GetUsers)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", GetUser)
			r.Delete("/", DeleteUser)
			r.Put("/", UpdateUser)
		})
	})
	r.Mount("/plans", plan.Router())
	r.Mount("/memberships", membership.Router())
	r.Mount("/payments", payment.Router())
	r.Mount("/products", product.Router())

	r.Mount("/tags", tag.Router())
	r.Mount("/categories", category.Router())

	r.Mount("/carts", cart.Router())
	r.Mount("/cart-items", cartitem.Router())

	r.Mount("/orders", order.Router())

	r.Route("/order-items", func(r chi.Router) {
		r.Post("/", CreateOrderItem)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", GetOrderItem)
			r.Delete("/", DeleteOrderItem)
			r.Put("/", UpdateOrderItem)
		})
	})

	// swagger docs
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	/* disable swagger in production */
	return r
}
