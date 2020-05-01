package membership

import (
	"encoding/json"
	"net/http"

	"github.com/factly/data-portal-server/model"
	"github.com/factly/data-portal-server/util"
)

// list - Get all memberships
// @Summary Show all memberships
// @Description Get all memberships
// @Tags Membership
// @ID get-all-memberships
// @Produce  json
// @Param limit query string false "limt per page"
// @Param page query string false "page number"
// @Success 200 {array} model.Membership
// @Router /memberships [get]
func list(w http.ResponseWriter, r *http.Request) {

	var memberships []model.Membership

	offset, limit := util.Paging(r.URL.Query())

	model.DB.Offset(offset).Limit(limit).Preload("User").Preload("Plan").Preload("Payment").Preload("Payment.Currency").Model(&model.Membership{}).Find(&memberships)

	json.NewEncoder(w).Encode(memberships)
}
