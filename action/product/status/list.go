package status

import (
	"net/http"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/data-portal-server/util"
	"github.com/factly/data-portal-server/util/render"
)

// list response
type paging struct {
	Total int            `json:"total"`
	Nodes []model.Status `json:"nodes"`
}

// list - Get all statuses
// @Summary Show all statuses
// @Description Get all statuses
// @Tags Status
// @ID get-all-statuses
// @Produce  json
// @Param limit query string false "limt per page"
// @Param page query string false "page number"
// @Success 200 {object} paging
// @Router /products/{product_id}/status [get]
func list(w http.ResponseWriter, r *http.Request) {

	result := paging{}

	offset, limit := util.Paging(r.URL.Query())

	model.DB.Model(&model.Status{}).Count(&result.Total).Offset(offset).Limit(limit).Find(&result.Nodes)

	render.JSON(w, http.StatusOK, result)
}
