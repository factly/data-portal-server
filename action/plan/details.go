package plan

import (
	"net/http"
	"strconv"

	"github.com/factly/mande-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// details - Get plan by id
// @Summary Show a plan by id
// @Description Get plan by ID
// @Tags Plan
// @ID get-plan-by-id
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param plan_id path string true "Plan ID"
// @Success 200 {object} model.Plan
// @Failure 400 {array} string
// @Router /plans/{plan_id} [get]
func details(w http.ResponseWriter, r *http.Request) {

	planID := chi.URLParam(r, "plan_id")
	id, err := strconv.Atoi(planID)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	result := &model.Plan{}
	result.ID = uint(id)

	err = model.DB.Model(&model.Plan{}).Preload("Currency").First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	if !result.AllProducts {
		err = model.DB.Model(&model.Plan{}).Preload("Currency").Preload("Catalogs").Preload("Catalogs.Products").Preload("Catalogs.Products.Currency").Preload("Catalogs.Products.Datasets").Preload("Catalogs.Products.Tags").First(&result).Error
	}

	renderx.JSON(w, http.StatusOK, result)
}
