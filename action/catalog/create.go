package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/factly/mande-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/meilisearchx"
	"github.com/factly/x/middlewarex"
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
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Catalog body catalog true "Catalog object"
// @Success 201 {object} model.Catalog
// @Failure 400 {array} string
// @Router /catalogs [post]
func create(w http.ResponseWriter, r *http.Request) {
	uID, err := middlewarex.GetUser(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	catalog := catalog{}
	result := model.Catalog{}
	result.Products = make([]model.Product, 0)

	err = json.NewDecoder(r.Body).Decode(&catalog)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(catalog)
	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	featuredMediumID := &catalog.FeaturedMediumID
	if catalog.FeaturedMediumID == 0 {
		featuredMediumID = nil
	}

	result = model.Catalog{
		Title:            catalog.Title,
		Slug:             catalog.Slug,
		Price:            catalog.Price,
		CurrencyID:       catalog.CurrencyID,
		Description:      catalog.Description,
		FeaturedMediumID: featuredMediumID,
		PublishedDate:    catalog.PublishedDate,
	}

	if len(catalog.ProductTitles) > 0 {
		model.DB.Model(&model.Product{}).Where("title IN ?", catalog.ProductTitles).Find(&result.Products)
	}

	if len(catalog.ProductIDs) > 0 {
		model.DB.Model(&model.Product{}).Where(catalog.ProductIDs).Find(&result.Products)
	}

	tx := model.DB.WithContext(context.WithValue(r.Context(), userContext, uID)).Begin()
	err = tx.Model(&model.Catalog{}).Create(&result).Error

	if err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	tx.Preload("FeaturedMedium").Preload("Currency").Preload("Products").Preload("Products.Currency").Preload("Products.FeaturedMedium").Preload("Products.Tags").Preload("Products.Datasets").First(&result)

	var meiliPublishedDate int64 = 0
	if result.PublishedDate != nil {
		meiliPublishedDate = result.PublishedDate.Unix()
	}
	// Insert into meili index
	meiliObj := map[string]interface{}{
		"id":             result.ID,
		"kind":           "catalog",
		"title":          result.Title,
		"slug":           result.Slug,
		"price":          result.Price,
		"currency_id":    result.CurrencyID,
		"description":    result.Description,
		"published_date": meiliPublishedDate,
		"product_ids":    catalog.ProductIDs,
	}

	err = meilisearchx.AddDocument("mande", meiliObj)
	if err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	tx.Commit()

	renderx.JSON(w, http.StatusCreated, result)
}
