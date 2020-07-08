package format

import (
	"net/http"
	"strconv"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// delete - Delete dataset format by id
// @Summary Delete a dataset format
// @Description Delete dataset format by ID
// @Tags Dataset Format
// @ID delete-dataset-format-by-id
// @Consume  json
// @Param dataset_id path string true "Dataset ID"
// @Param format_id path string true "Dataset Format ID"
// @Success 200
// @Failure 400 {array} string
// @Router /datasets/{dataset_id}/format/{format_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

	datasetID := chi.URLParam(r, "dataset_id")
	id, err := strconv.Atoi(datasetID)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	formatID := chi.URLParam(r, "format_id")
	fID, err := strconv.Atoi(formatID)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	result := &model.DatasetFormat{}
	result.DatasetID = uint(id)
	result.FormatID = uint(fID)

	// check record exists or not
	err = model.DB.First(&result).Error
	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	model.DB.Delete(&result)

	renderx.JSON(w, http.StatusOK, nil)
}
