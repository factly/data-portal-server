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

	r.Get("/", getPayments)    // GET /payments - return list of payments
	r.Post("/", createPayment) // POST /payments - create a new payment and persist it

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", getPaymentByID)   // GET /payments/{id} - read a single payment by :id
		r.Put("/", updatePayment)    // PUT /payments/{id} - update a single payment by :id
		r.Delete("/", deletePayment) // DELETE /payments/{id} - delete a single payment by :id
	})

	return r
}
