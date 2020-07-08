package cart

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// update - Update cart by id
// @Summary Update a cart by id
// @Description Update cart by ID
// @Tags Cart
// @ID update-cart-by-id
// @Produce json
// @Consume json
// @Param cart_id path string true "Cart ID"
// @Param Cart body cart false "Cart"
// @Success 200 {object} model.Cart
// @Failure 400 {array} string
// @Router /carts/{cart_id} [put]
func update(w http.ResponseWriter, r *http.Request) {
	cartID := chi.URLParam(r, "cart_id")
	id, err := strconv.Atoi(cartID)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	cart := &cart{}

	json.NewDecoder(r.Body).Decode(&cart)

	result := &model.Cart{}
	result.ID = uint(id)

	model.DB.Model(&result).Updates(model.Cart{
		Status: cart.Status,
		UserID: cart.UserID,
	}).First(&result)

	renderx.JSON(w, http.StatusOK, result)
}
