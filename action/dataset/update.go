package dataset

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/data-portal-server/validation"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// update - Update dataset by id
// @Summary Update a dataset by id
// @Description Update dataset by ID
// @Tags Dataset
// @ID update-dataset-by-id
// @Produce json
// @Consume json
// @Param dataset_id path string true "Dataset ID"
// @Param Dataset body dataset false "Dataset"
// @Success 200 {object} model.Dataset
// @Failure 400 {array} string
// @Router /datasets/{dataset_id} [put]
func update(w http.ResponseWriter, r *http.Request) {

	datasetID := chi.URLParam(r, "dataset_id")
	id, err := strconv.Atoi(datasetID)

	if err != nil {
		validation.InvalidID(w, r)
		return
	}

	dataset := &dataset{}
	formats := []model.DatasetFormat{}
	result := &datasetData{}
	result.ID = uint(id)

	json.NewDecoder(r.Body).Decode(&dataset)

	model.DB.Model(&result).Updates(model.Dataset{
		Title:            dataset.Title,
		Description:      dataset.Description,
		Source:           dataset.Source,
		Frequency:        dataset.Frequency,
		TemporalCoverage: dataset.TemporalCoverage,
		Granularity:      dataset.Granularity,
		ContactName:      dataset.ContactName,
		ContactEmail:     dataset.ContactEmail,
		License:          dataset.License,
		DataStandard:     dataset.DataStandard,
		RelatedArticles:  dataset.RelatedArticles,
		TimeSaved:        dataset.TimeSaved,
	}).First(&result)

	model.DB.Model(&model.DatasetFormat{}).Where(&model.DatasetFormat{
		DatasetID: uint(id),
	}).Preload("Format").Find(&formats)

	for _, f := range formats {
		result.Formats = append(result.Formats, f.Format)
	}

	renderx.JSON(w, http.StatusOK, result)
}