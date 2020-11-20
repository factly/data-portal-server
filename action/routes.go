package action

import (
	"fmt"
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
	_ "github.com/factly/data-portal-server/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/factly/x/loggerx"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

// GetCommonRouter returns router with common middleware and settings
func GetCommonRouter() chi.Router {
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

	return r
}

// RegisterUserRoutes - register user routes
func RegisterUserRoutes() http.Handler {

	r := GetCommonRouter()

	r.Mount("/currencies", currency.UserRouter())
	r.Mount("/plans", plan.UserRouter())
	r.Mount("/memberships", membership.UserRouter())
	r.Mount("/payments", payment.UserRouter())
	r.Mount("/products", product.UserRouter())
	r.Mount("/tags", tag.UserRouter())
	r.Mount("/formats", format.UserRouter())
	r.Mount("/catalogs", catalog.UserRouter())
	r.Mount("/cartitems", cart.UserRouter())
	r.Mount("/orders", order.UserRouter())
	r.Mount("/datasets", dataset.UserRouter())
	r.Mount("/media", medium.UserRouter())
	r.Mount("/search", search.Router())

	return r
}

// RegisterAdminRoutes - register admin routes
func RegisterAdminRoutes() http.Handler {

	r := GetCommonRouter()

	r.Mount("/currencies", currency.AdminRouter())
	r.Mount("/plans", plan.AdminRouter())
	r.Mount("/memberships", membership.AdminRouter())
	r.Mount("/payments", payment.AdminRouter())
	r.Mount("/products", product.AdminRouter())
	r.Mount("/tags", tag.AdminRouter())
	r.Mount("/formats", format.AdminRouter())
	r.Mount("/catalogs", catalog.AdminRouter())
	r.Mount("/cartitems", cart.AdminRouter())
	r.Mount("/orders", order.AdminRouter())
	r.Mount("/datasets", dataset.AdminRouter())
	r.Mount("/media", medium.AdminRouter())
	r.Mount("/search", search.Router())

	if viper.IsSet("mode") && viper.GetString("mode") == "development" {
		r.Get("/swagger/*", httpSwagger.WrapHandler)
		fmt.Println("Admin Swagger @ http://localhost:7721/swagger/index.html")
	}

	return r
}
