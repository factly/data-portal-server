package order

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
	"github.com/go-chi/chi"
)

// update - Update orders by id
// @Summary Update a orders by id
// @Description Update orders by ID
// @Tags Order
// @ID update-orders-by-id
// @Produce json
// @Consume json
// @Param order_id path string true "Order ID"
// @Param Order body order false "Order"
// @Success 200 {object} model.Order
// @Failure 400 {array} string
// @Router /orders/{order_id} [put]
func update(w http.ResponseWriter, r *http.Request) {

	orderID := chi.URLParam(r, "order_id")
	id, err := strconv.Atoi(orderID)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	order := &order{}

	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(order)
	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	result := &model.Order{}
	result.ID = uint(id)

	model.DB.Model(&result).Updates(model.Order{
		UserID:    order.UserID,
		PaymentID: order.PaymentID,
		Status:    order.Status,
		CartID:    order.CartID,
	}).Preload("Payment").Preload("Payment.Currency").Preload("Cart").First(&result)

	renderx.JSON(w, http.StatusOK, result)
}
