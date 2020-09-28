package order

import (
	"github.com/go-chi/chi"
)

// UserRouter - Group of order router
func UserRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", userlist) // GET /orders - return list of orders
	r.Post("/", create)  // POST /orders - create a new order and persist it

	r.Route("/{order_id}", func(r chi.Router) {
		r.Get("/", details)   // GET /orders/{order_id} - read a single order by :order_id
		r.Delete("/", delete) // DELETE /orders/{order_id} - delete a single order by :order_id
	})

	return r
}

// AdminRouter - Group of order router
func AdminRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", adminlist) // GET /orders - return list of orders

	r.Route("/{order_id}", func(r chi.Router) {
		r.Get("/", details) // GET /orders/{order_id} - read a single order by :order_id
	})

	return r
}
