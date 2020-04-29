package currency

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/factly/data-portal-api/model"
)

// getCurrencies - Get all currencies
// @Summary Show all currencies
// @Description Get all currencies
// @Tags Currency
// @ID get-all-currencies
// @Produce  json
// @Param limit query string false "limt per page"
// @Param page query string false "page number"
// @Success 200 {array} model.Currency
// @Router /currencies [get]
func getCurrencies(w http.ResponseWriter, r *http.Request) {

	var currencies []model.Currency
	p := r.URL.Query().Get("page")
	pg, _ := strconv.Atoi(p) // pg contains page number
	l := r.URL.Query().Get("limit")
	li, _ := strconv.Atoi(l) // li contains perPage number

	offset := 0 // no. of records to skip
	limit := 5  // limt

	if li > 0 && li <= 10 {
		limit = li
	}

	if pg > 1 {
		offset = (pg - 1) * limit
	}

	model.DB.Offset(offset).Limit(limit).Model(&model.Currency{}).Find(&currencies)

	json.NewEncoder(w).Encode(currencies)
}