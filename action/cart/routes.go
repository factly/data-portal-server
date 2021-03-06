package cart

import (
	"github.com/factly/mande-server/model"
	"github.com/go-chi/chi"
)

// CartItem request body
type cartitem struct {
	Status       string `json:"status" validate:"required"`
	ProductID    uint   `json:"product_id" validate:"required"`
	MembershipID uint   `json:"membership_id"`
}

var userContext model.ContextKey = "cartitem_user"

// UserRouter - Group of user cart router
func UserRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", userList) // GET /carts - return list of cart items
	r.Post("/", create)  // POST /carts - add a new cart item

	r.Route("/{cartitem_id}", func(r chi.Router) {
		r.Get("/", userDetails) // GET /carts/{cartitem_id} - get a single cart item
		r.Delete("/", delete)   // DELETE /carts/{cartitem_id} - delete a cart item entry
	})

	return r
}

// AdminRouter - Group of user cart router
func AdminRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", adminList)
	r.Route("/{cartitem_id}", func(r chi.Router) {
		r.Get("/", adminDetails) // GET /carts/{cartitem_id} - get a single cart item
	})

	return r
}
