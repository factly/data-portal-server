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

// delete - Delete catalog by id
// @Summary Delete a catalog
// @Description Delete catalog by ID
// @Tags Catalog
// @ID delete-catalog-by-id
// @Consume  json
// @Param catalog_id path string true "Catalog ID"
// @Success 200
// @Failure 400 {array} string
// @Router /catalogs/{catalog_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

	catalogID := chi.URLParam(r, "catalog_id")
	id, err := strconv.Atoi(catalogID)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	result := &model.Catalog{}
	result.ID = uint(id)

	// check record exists or not
	err = model.DB.First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}
	model.DB.Delete(&result)

	renderx.JSON(w, http.StatusOK, nil)
}
