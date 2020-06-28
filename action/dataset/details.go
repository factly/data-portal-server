package dataset

import (
	"net/http"
	"strconv"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/data-portal-server/validation"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// details - Get dataset by id
// @Summary Show a dataset by id
// @Description Get dataset by ID
// @Tags Dataset
// @ID get-dataset-by-id
// @Produce  json
// @Param dataset_id path string true "Dataset ID"
// @Success 200 {object} model.Dataset
// @Failure 400 {array} string
// @Router /datasets/{dataset_id} [get]
func details(w http.ResponseWriter, r *http.Request) {

	datasetID := chi.URLParam(r, "dataset_id")
	id, err := strconv.Atoi(datasetID)

	if err != nil {
		validation.InvalidID(w, r)
		return
	}

	result := &datasetData{}
	result.ID = uint(id)
	result.Formats = make([]model.DatasetFormat, 0)

	err = model.DB.Model(&model.Dataset{}).Preload("FeaturedMedia").First(&result.Dataset).Error

	if err != nil {
		validation.RecordNotFound(w, r)
		return
	}

	model.DB.Model(&model.DatasetFormat{}).Where(&model.DatasetFormat{
		DatasetID: uint(id),
	}).Preload("Format").Find(&result.Formats)

	renderx.JSON(w, http.StatusOK, result)
}
