package catalog

import (
	"net/http"
	"strconv"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// details - Get catalog by id
// @Summary Show a catalog by id
// @Description Get catalog by ID
// @Tags Catalog
// @ID get-catalog-by-id
// @Produce  json
// @Param catalog_id path string true "Catalog ID"
// @Success 200 {object} model.Catalog
// @Failure 400 {array} string
// @Router /catalogs/{catalog_id} [get]
func details(w http.ResponseWriter, r *http.Request) {

	products := []model.CatalogProduct{}
	catalogID := chi.URLParam(r, "catalog_id")
	id, err := strconv.Atoi(catalogID)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	result := &catalogData{}
	result.ID = uint(id)
	result.Products = make([]model.Product, 0)

	err = model.DB.Model(&model.Catalog{}).Preload("FeaturedMedium").First(&result.Catalog).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	model.DB.Model(&model.CatalogProduct{}).Where(&model.CatalogProduct{
		CatalogID: uint(id),
	}).Preload("Product").Find(&products)

	for _, p := range products {
		result.Products = append(result.Products, p.Product)
	}

	renderx.JSON(w, http.StatusOK, result)
}
