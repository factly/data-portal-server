package cart

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
)

// create - create cart
// @Summary Create cart
// @Description create cart
// @Tags Cart
// @ID add-cart
// @Consume json
// @Produce  json
// @Param Cart body cart true "Cart object"
// @Success 201 {object} model.Cart
// @Failure 400 {array} string
// @Router /carts [post]
func create(w http.ResponseWriter, r *http.Request) {

	cart := &cart{}

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(cart)
	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	result := &model.Cart{
		Status: cart.Status,
		UserID: cart.UserID,
	}

	model.DB.Model(&model.Product{}).Where(cart.ProductIDs).Find(&result.Products)

	err = model.DB.Model(&model.Cart{}).Set("gorm:association_autoupdate", false).Create(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	model.DB.Preload("Products").Preload("Products.Currency").Preload("Products.FeaturedMedium").Preload("Products.Tags").Preload("Products.Datasets").First(&result)

	renderx.JSON(w, http.StatusCreated, result)
}
