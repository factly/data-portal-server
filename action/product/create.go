package product

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

// create - Create product
// @Summary Create product
// @Description Create product
// @Tags Product
// @ID add-product
// @Consume json
// @Produce  json
// @Param Product body product true "Product object"
// @Success 201 {object} model.Product
// @Failure 400 {array} string
// @Router /products [post]
func create(w http.ResponseWriter, r *http.Request) {

	product := &product{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(product)
	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	result := model.Product{}
	result.Tags = make([]model.Tag, 0)
	result.Datasets = make([]model.Dataset, 0)
	result = model.Product{
		Title:            product.Title,
		Slug:             product.Slug,
		Price:            product.Price,
		Status:           product.Status,
		CurrencyID:       product.CurrencyID,
		FeaturedMediumID: product.FeaturedMediumID,
	}

	model.DB.Model(&model.Tag{}).Where(product.TagIDs).Find(&result.Tags)
	model.DB.Model(&model.Dataset{}).Where(product.DatasetIDs).Find(&result.Datasets)

	err = model.DB.Model(&model.Product{}).Set("gorm:association_autoupdate", false).Create(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	model.DB.Preload("Currency").Preload("FeaturedMedium").Preload("Tags").Preload("Datasets").First(&result)

	renderx.JSON(w, http.StatusCreated, result)
}
