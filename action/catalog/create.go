package catalog

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

// create - create catalog
// @Summary Create catalog
// @Description create catalog
// @Tags Catalog
// @ID add-catalog
// @Consume json
// @Produce  json
// @Param Catalog body catalog true "Catalog object"
// @Success 201 {object} model.Catalog
// @Failure 400 {array} string
// @Router /catalogs [post]
func create(w http.ResponseWriter, r *http.Request) {

	catalog := catalog{}
	result := catalogData{}
	result.Products = make([]model.Product, 0)

	json.NewDecoder(r.Body).Decode(&catalog)

	validationError := validationx.Check(catalog)
	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	result.Catalog = model.Catalog{
		Title:            catalog.Title,
		Description:      catalog.Description,
		FeaturedMediumID: catalog.FeaturedMediumID,
		Price:            catalog.Price,
		PublishedDate:    catalog.PublishedDate,
	}

	err := model.DB.Model(&model.Catalog{}).Create(&result.Catalog).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	model.DB.Preload("FeaturedMedium").First(&result.Catalog)

	for _, id := range catalog.ProductIDs {
		catalogProduct := &model.CatalogProduct{}

		catalogProduct.ProductID = uint(id)
		catalogProduct.CatalogID = result.ID

		err = model.DB.Model(&model.CatalogProduct{}).Create(&catalogProduct).Error

		if err != nil {
			errorx.Render(w, errorx.Parser(errorx.DBError()))
			return
		}
		model.DB.Model(&model.CatalogProduct{}).Preload("Product").First(&catalogProduct)
		result.Products = append(result.Products, catalogProduct.Product)
	}

	renderx.JSON(w, http.StatusCreated, result)
}
