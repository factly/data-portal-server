package dataset

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

// create - create dataset
// @Summary Create dataset
// @Description create dataset
// @Tags Dataset
// @ID add-dataset
// @Consume json
// @Produce  json
// @Param Dataset body dataset true "Dataset object"
// @Success 201 {object} model.Dataset
// @Failure 400 {array} string
// @Router /datasets [post]
func create(w http.ResponseWriter, r *http.Request) {

	dataset := dataset{}

	err := json.NewDecoder(r.Body).Decode(&dataset)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(dataset)
	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	result := &datasetData{}
	result.Tags = make([]model.Tag, 0)
	result.Dataset = model.Dataset{
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
		FeaturedMediumID: dataset.FeaturedMediumID,
		Price:            dataset.Price,
		CurrencyID:       dataset.CurrencyID,
	}

	model.DB.Model(&model.Tag{}).Where(dataset.TagIDs).Find(&result.Tags)

	err = model.DB.Model(&model.Dataset{}).Set("gorm:association_autoupdate", false).Create(&result.Dataset).Error
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	model.DB.Preload("FeaturedMedium").Preload("Currency").Preload("Tags").First(&result.Dataset)

	renderx.JSON(w, http.StatusCreated, result)
}
