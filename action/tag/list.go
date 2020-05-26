package tag

import (
	"net/http"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/data-portal-server/util"
	"github.com/factly/data-portal-server/util/render"
)

// list response
type paging struct {
	Total int         `json:"total"`
	Nodes []model.Tag `json:"nodes"`
}

// @Summary Show all tags
// @Description Get all tags
// @Tags Tag
// @ID get-all-tags
// @Produce  json
// @Param limit query string false "limt per page"
// @Param page query string false "page number"
// @Success 200 {object} paging
// @Router /tags [get]
func list(w http.ResponseWriter, r *http.Request) {

	result := paging{}

	offset, limit := util.Paging(r.URL.Query())

	model.DB.Model(&model.Tag{}).Count(&result.Total).Offset(offset).Limit(limit).Find(&result.Nodes)

	render.JSON(w, http.StatusOK, result)
}
