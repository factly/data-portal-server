package item

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/data-portal-server/util/render"
	"github.com/factly/data-portal-server/validation"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

// create - create order items
// @Summary Create order items
// @Description create order items
// @Tags OrderItem
// @ID add-order-item
// @Consume json
// @Produce  json
// @Param order_id path string true "Order ID"
// @Param OrderItem body orderItem true "Order item object"
// @Success 201 {object} model.OrderItem
// @Router /orders/{order_id}/items [post]
func create(w http.ResponseWriter, r *http.Request) {

	orderID := chi.URLParam(r, "order_id")
	id, _ := strconv.Atoi(orderID)

	orderItem := &orderItem{}
	result := &model.OrderItem{}

	json.NewDecoder(r.Body).Decode(&orderItem)

	validate := validator.New()
	err := validate.StructExcept(orderItem, "Product", "Order")
	if err != nil {
		msg := err.Error()
		validation.ValidErrors(w, r, msg)
		return
	}

	result.OrderID = uint(id)
	result.ExtraInfo = orderItem.ExtraInfo
	result.ProductID = orderItem.ProductID

	err = model.DB.Model(&model.OrderItem{}).Create(&result).Error

	if err != nil {
		log.Fatal(err)
	}
	model.DB.Model(&model.OrderItem{}).Preload("Product").Preload("Product.Status").Preload("Product.ProductType").Preload("Product.Currency").Preload("Order").First(&result)

	render.JSON(w, http.StatusCreated, result)
}
