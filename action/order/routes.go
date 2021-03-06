package order

import (
	"github.com/factly/mande-server/model"
	"github.com/go-chi/chi"
)

// UserRouter - Group of order router
func UserRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", userList) // GET /orders - return list of orders
	r.Post("/", create)  // POST /orders - create a new order and persist it

	r.Route("/{order_id}", func(r chi.Router) {
		r.Get("/", userDetails) // GET /orders/{order_id} - read a single order by :order_id
		r.Delete("/", delete)   // DELETE /orders/{order_id} - delete a single order by :order_id
	})

	return r
}

var userContext model.ContextKey = "order_user"

// AdminRouter - Group of order router
func AdminRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", adminList) // GET /orders - return list of orders

	r.Route("/{order_id}", func(r chi.Router) {
		r.Get("/", adminDetails) // GET /orders/{order_id} - read a single order by :order_id
	})

	return r
}
