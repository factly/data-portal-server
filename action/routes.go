package action

import (
	"net/http"

	"github.com/factly/data-portal-server/action/cart"
	"github.com/factly/data-portal-server/action/catalog"
	"github.com/factly/data-portal-server/action/currency"
	"github.com/factly/data-portal-server/action/dataset"
	"github.com/factly/data-portal-server/action/format"
	"github.com/factly/data-portal-server/action/medium"
	"github.com/factly/data-portal-server/action/membership"
	"github.com/factly/data-portal-server/action/order"
	"github.com/factly/data-portal-server/action/payment"
	"github.com/factly/data-portal-server/action/plan"
	"github.com/factly/data-portal-server/action/product"
	"github.com/factly/data-portal-server/action/search"
	"github.com/factly/data-portal-server/action/tag"
	"github.com/factly/data-portal-server/action/user"
	_ "github.com/factly/data-portal-server/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/factly/x/loggerx"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterRoutes - register routes
func RegisterRoutes() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(loggerx.Init())
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/ping"))

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

	/* disable swagger in production */
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Mount("/currencies", currency.Router())
	r.Mount("/users", user.Router())
	r.Mount("/plans", plan.Router())
	r.Mount("/memberships", membership.Router())
	r.Mount("/payments", payment.Router())
	r.Mount("/products", product.Router())
	r.Mount("/tags", tag.Router())
	r.Mount("/formats", format.Router())
	r.Mount("/catalogs", catalog.Router())
	r.Mount("/carts", cart.Router())
	r.Mount("/orders", order.Router())
	r.Mount("/datasets", dataset.Router())
	r.Mount("/media", medium.Router())
	r.Mount("/search", search.Router())

	return r
}
