package payment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// update - Update payment by id
// @Summary Update a payment by id
// @Description Update payment by ID
// @Tags Payment
// @ID update-payment-by-id
// @Produce json
// @Consume json
// @Param payment_id path string true "Payment ID"
// @Param Payment body payment false "Payment"
// @Success 200 {object} model.Payment
// @Failure 400 {array} string
// @Router /payments/{payment_id} [put]
func update(w http.ResponseWriter, r *http.Request) {

	paymentID := chi.URLParam(r, "payment_id")
	id, err := strconv.Atoi(paymentID)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	payment := &payment{}

	json.NewDecoder(r.Body).Decode(&payment)

	result := &model.Payment{}
	result.ID = uint(id)

	model.DB.Model(&result).Updates(&model.Payment{
		Amount:     payment.Amount,
		Gateway:    payment.Gateway,
		Status:     payment.Status,
		CurrencyID: payment.CurrencyID,
	}).First(&result).Preload("Currency").Find(&result.Currency)

	renderx.JSON(w, http.StatusOK, result)
}
