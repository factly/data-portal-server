package payment

import "github.com/go-chi/chi"

// payment request body
type payment struct {
	Amount     int    `json:"amount"`
	Gateway    string `json:"gateway"`
	CurrencyID uint   `json:"currency_id"`
	Status     string `json:"status"`
}

// Router - Group of payment router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)    // GET /payments - return list of payments
	r.Post("/", create) // POST /payments - create a new payment and persist it

	r.Route("/{payment_id}", func(r chi.Router) {
		r.Get("/", details)   // GET /payments/{payment_id} - read a single payment by :payment_id
		r.Put("/", update)    // PUT /payments/{payment_id} - update a single payment by :payment_id
		r.Delete("/", delete) // DELETE /payments/{payment_id} - delete a single payment by :payment_id
	})

	return r
}
